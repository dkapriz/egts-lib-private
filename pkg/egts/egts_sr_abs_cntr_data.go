package egts

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// SrAbsCntrData структура под записи типа EGTS_SR_ABS_CNTR_DATA, которая применяется
// абонентским терминалом для передачи на аппаратно-программный комплекс данных о
// состоянии одного счетного входа
type SrAbsCntrData struct {
	CounterNumber uint8  `json:"CN"`
	CounterValue  uint32 `json:"CNV"`
}

// Decode разбирает байты в структуры секции
func (e *SrAbsCntrData) Decode(content []byte, _ byte) error {
	var (
		err error
	)
	buf := bytes.NewReader(content)

	if e.CounterNumber, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("не удалось получить номер счетного входа: %v", err)
	}

	tmpBuf := make([]byte, 3)
	if _, err = buf.Read(tmpBuf); err != nil {
		return fmt.Errorf("не удалось получить значение показаний счетного входа: %v", err)
	}

	counterVal := append(tmpBuf, 0x00)
	e.CounterValue = binary.LittleEndian.Uint32(counterVal)

	return err
}

// Encode преобразовывает под запись в набор байт
func (e *SrAbsCntrData) Encode(_ byte) ([]byte, error) {
	var (
		err    error
		result []byte
	)
	buf := new(bytes.Buffer)

	if err = buf.WriteByte(e.CounterNumber); err != nil {
		return result, fmt.Errorf("не удалось записать номер счетного входа: %v", err)
	}

	counterVal := make([]byte, 4)
	binary.LittleEndian.PutUint32(counterVal, e.CounterValue)
	if _, err = buf.Write(counterVal[:3]); err != nil {
		return result, fmt.Errorf("не удалось записать значение показаний счетного входа: %v", err)
	}

	result = buf.Bytes()
	return result, err
}

// Length получает длину закодированной под записи
func (e *SrAbsCntrData) Length(protocolVersion byte) uint16 {
	var result uint16

	if recBytes, err := e.Encode(protocolVersion); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
