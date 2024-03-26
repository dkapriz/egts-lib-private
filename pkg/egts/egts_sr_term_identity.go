package egts

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

// SrTermIdentity структура под записи типа EGTS_SR_TERM_IDENTITY, которая используется АС при запросе
// авторизации на телематическую платформу и содержит учетные данные АС.
type SrTermIdentity struct {
	TerminalIdentifier                 uint64 `json:"TID"` //Protocol version 2 (увеличена длина с 4 до 8 байт)
	MNE                                string `json:"MNE"`
	BSE                                string `json:"BSE"`
	NIDE                               string `json:"NIDE"`
	SSRA                               string `json:"SSRA"`
	LNGCE                              string `json:"LNGCE"`
	IMSIE                              string `json:"IMSIE"`
	IMEIE                              string `json:"IMEIE"`
	HDIDE                              string `json:"HDIDE"`
	HomeDispatcherIdentifier           uint16 `json:"HDID"`
	IMEI                               string `json:"IMEI"`
	IMSI                               string `json:"IMSI"`
	LanguageCode                       string `json:"LNGC"`
	NetworkIdentifier                  []byte `json:"NID"`
	BufferSize                         uint16 `json:"BS"`
	MobileNumber                       string `json:"MSISDN"`
	ServiceSupportLevelProtocolVersion string `json:"SSLPV"` //Protocol version 2
}

// Decode разбирает байты в структуру под записи
func (e *SrTermIdentity) Decode(content []byte, protocolVersion byte) error {
	var (
		err   error
		flags byte
	)
	buf := bytes.NewReader(content)
	tmpBuf := make([]byte, 4)

	if protocolVersion == 1 {
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("не удалось получить идентификатор терминал при авторизации")
		}
		e.TerminalIdentifier = uint64(binary.LittleEndian.Uint32(tmpBuf))
	}
	if protocolVersion >= 2 {
		tmpBuf = make([]byte, 8)
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("не удалось получить идентификатор терминал при авторизации")
		}
		e.TerminalIdentifier = binary.LittleEndian.Uint64(tmpBuf)
	}

	tmpBuf = make([]byte, 4)
	if flags, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("не удалось считать байт флагов term identify: %v", err)
	}
	flagBits := fmt.Sprintf("%08b", flags)
	e.MNE = flagBits[:1]
	e.BSE = flagBits[1:2]
	e.NIDE = flagBits[2:3]
	e.SSRA = flagBits[3:4]
	e.LNGCE = flagBits[4:5]
	e.IMSIE = flagBits[5:6]
	e.IMEIE = flagBits[6:7]
	e.HDIDE = flagBits[7:]

	if e.HDIDE == "1" {
		tmpBuf = make([]byte, 2)
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("не удалось получить идентификатор «домашней» телематической платформы " +
				"при авторизации")
		}
		e.HomeDispatcherIdentifier = binary.LittleEndian.Uint16(tmpBuf)

	}

	if e.IMEIE == "1" {
		tmpBuf = make([]byte, 15)
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("не удалось получить IMEI при авторизации")
		}
		e.IMEI = string(tmpBuf)
	}

	if e.IMSIE == "1" {
		tmpBuf = make([]byte, 16)
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("не удалось получить IMSI при авторизации")
		}
		e.IMSI = string(tmpBuf)
	}

	if e.LNGCE == "1" {
		tmpBuf = make([]byte, 3)
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("не удалось получить код языка при авторизации")
		}
		e.LanguageCode = string(tmpBuf)
	}

	if e.NIDE == "1" {
		e.NetworkIdentifier = make([]byte, 3)
		if _, err = buf.Read(e.NetworkIdentifier); err != nil {
			return fmt.Errorf("не удалось получить код идентификатор сети оператора при авторизации")
		}
	}

	if e.BSE == "1" {
		tmpBuf = make([]byte, 2)
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("не удалось получить максимальный размер буфера при авторизации")
		}
		e.BufferSize = binary.LittleEndian.Uint16(tmpBuf)
	}

	if e.MNE == "1" {
		tmpBuf = make([]byte, 15)
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("не удалось получить телефонный номер мобильного абонента")
		}
		e.MobileNumber = string(tmpBuf)
	}

	//Protocol version 2
	if protocolVersion >= 2 {
		tmpBuf = make([]byte, 2)
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("не удалось получить номер версии протокола уровня поддержки услуг")
		}
		e.ServiceSupportLevelProtocolVersion = string(tmpBuf)
	}

	return err
}

// Encode преобразовывает под запись в набор байт
func (e *SrTermIdentity) Encode(protocolVersion byte) ([]byte, error) {
	var (
		result []byte
		flags  uint64
		err    error
	)
	buf := new(bytes.Buffer)

	if protocolVersion == 1 {
		if err = binary.Write(buf, binary.LittleEndian, uint32(e.TerminalIdentifier)); err != nil {
			return result, fmt.Errorf("не удалось записать идентификатор терминала при авторизации")
		}
	} else {
		//Protocol version 2
		if err = binary.Write(buf, binary.LittleEndian, e.TerminalIdentifier); err != nil {
			return result, fmt.Errorf("не удалось записать идентификатор терминала при авторизации")
		}
	}

	flags, err = strconv.ParseUint(e.MNE+e.BSE+e.NIDE+e.SSRA+e.LNGCE+e.IMSIE+e.IMEIE+e.HDIDE, 2, 8)
	if err = buf.WriteByte(uint8(flags)); err != nil {
		return result, fmt.Errorf("не удалось записать байт флагов term identify: %v", err)
	}

	if e.HDIDE == "1" {
		if err = binary.Write(buf, binary.LittleEndian, e.HomeDispatcherIdentifier); err != nil {
			return result, fmt.Errorf("не удалось записать идентификатор «домашней» телематической " +
				"платформы при авторизации")
		}
	}

	if e.IMEIE == "1" {
		if _, err = buf.Write([]byte(e.IMEI)); err != nil {
			return result, fmt.Errorf("не удалось записать IMEI при авторизации")
		}
	}

	if e.IMSIE == "1" {
		if _, err = buf.Write([]byte(e.IMSI)); err != nil {
			return result, fmt.Errorf("не удалось записать IMSI при авторизации")
		}
	}

	if e.LNGCE == "1" {
		if _, err = buf.Write([]byte(e.LanguageCode)); err != nil {
			return result, fmt.Errorf("не удалось записать IMSI при авторизации")
		}
	}

	if e.NIDE == "1" {
		if _, err = buf.Write(e.NetworkIdentifier); err != nil {
			return result, fmt.Errorf("не удалось записать код идентификатор сети оператора при авторизации")
		}
	}

	if e.BSE == "1" {
		if err = binary.Write(buf, binary.LittleEndian, e.BufferSize); err != nil {
			return result, fmt.Errorf("не удалось записать максимальный размер буфера при авторизации")
		}
	}

	if e.MNE == "1" {
		if _, err = buf.Write([]byte(e.MobileNumber)); err != nil {
			return result, fmt.Errorf("не удалось записать телефонный номер мобильного абонента")
		}
	}

	if protocolVersion >= 2 {
		if _, err = buf.Write([]byte(e.ServiceSupportLevelProtocolVersion)); err != nil {
			return result, fmt.Errorf("не удалось записать номер версии протокола уровня поддержки услуг")
		}
	}

	result = buf.Bytes()
	return result, err
}

// Length получает длину закодированной под записи
func (e *SrTermIdentity) Length(protocolVersion byte) uint16 {
	var result uint16

	if recBytes, err := e.Encode(protocolVersion); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
