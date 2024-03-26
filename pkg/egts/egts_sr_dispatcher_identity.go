package egts

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// SrDispatcherIdentity структура под записи типа EGTS_SR_DISPATCHER_IDENTITY, которая используется
// только авторизуемой ТП при запросе авторизации на авторизующей ТП и содержит учетные данные
// авторизуемой АСН
type SrDispatcherIdentity struct {
	DispatcherType                     uint8  `json:"DT"`
	DispatcherID                       uint32 `json:"DID"`
	TerminalIdentifier                 uint64 `json:"TID"`   //Protocol version 2
	ServiceSupportLevelProtocolVersion string `json:"SSLPV"` //Protocol version 2
	Description                        string `json:"DSCR"`
}

// Decode разбирает байты в структуру под записи
func (d *SrDispatcherIdentity) Decode(content []byte, protocolVersion byte) error {
	var err error

	buf := bytes.NewBuffer(content)

	if d.DispatcherType, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("не удалось получить тип диспетчера: %v", err)
	}

	tmpIntBuf := make([]byte, 4)
	if _, err = buf.Read(tmpIntBuf); err != nil {
		return fmt.Errorf("не удалось получить уникальный идентификатор диспетчера: %v", err)
	}
	d.DispatcherID = binary.LittleEndian.Uint32(tmpIntBuf)

	if protocolVersion >= 2 {
		tmpIntBuf = make([]byte, 8)
		if _, err = buf.Read(tmpIntBuf); err != nil {
			return fmt.Errorf("не удалось получить уникальный идентификатор, назначаемый "+
				"при программировании УСВ: %v", err)
		}
		d.TerminalIdentifier = binary.LittleEndian.Uint64(tmpIntBuf)

		tmpIntBuf = make([]byte, 2)
		if _, err = buf.Read(tmpIntBuf); err != nil {
			return fmt.Errorf("не удалось получить номер версии ППУ авторизуемого УСВ")
		}
		d.ServiceSupportLevelProtocolVersion = string(tmpIntBuf)
	}

	d.Description = buf.String()

	return err
}

// Encode преобразовывает под запись в набор байт
func (d *SrDispatcherIdentity) Encode(protocolVersion byte) ([]byte, error) {
	var (
		result []byte
		err    error
	)

	buf := new(bytes.Buffer)

	if err = buf.WriteByte(d.DispatcherType); err != nil {
		return result, fmt.Errorf("не удалось записать тип диспетчера: %v", err)
	}

	if err = binary.Write(buf, binary.LittleEndian, d.DispatcherID); err != nil {
		return result, fmt.Errorf("не удалось записать уникальный идентификатор диспетчера: %v", err)
	}

	if protocolVersion >= 2 {
		if err = binary.Write(buf, binary.LittleEndian, d.TerminalIdentifier); err != nil {
			return result, fmt.Errorf("не удалось записать уникальный идентификатор, назначаемый "+
				"при программировании УСВ: %v", err)
		}

		if _, err = buf.Write([]byte(d.ServiceSupportLevelProtocolVersion)); err != nil {
			return result, fmt.Errorf("не удалось записать номер версии ППУ авторизуемого УСВ")
		}
	}

	if _, err = buf.WriteString(d.Description); err != nil {
		return result, fmt.Errorf("не удалось записать уникальный краткое описание: %v", err)
	}

	return buf.Bytes(), err
}

// Length получает длину закодированной под записи
func (d *SrDispatcherIdentity) Length(protocolVersion byte) uint16 {
	var result uint16

	if recBytes, err := d.Encode(protocolVersion); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
