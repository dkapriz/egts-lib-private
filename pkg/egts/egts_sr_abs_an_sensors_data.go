package egts

import (
	"encoding/binary"
	"errors"
)

// SrAbsAnSensData структура под записи типа EGTS_SR_ABS_AN_SENS_DATA, которая применяется абонентским
// терминалом для передачи данных о состоянии одного аналогового входа
type SrAbsAnSensData struct {
	SensorNumber uint8  `json:"SensorNumber"`
	Value        uint32 `json:"Value"`
}

// Decode разбирает байты в структуру под записи
func (e *SrAbsAnSensData) Decode(content []byte, protocolVersion byte) error {
	if len(content) < int(e.Length(protocolVersion)) {
		return errors.New("некорректный размер данных")
	}
	e.SensorNumber = content[0]
	e.Value = binary.LittleEndian.Uint32(content) >> 8
	return nil
}

// Encode преобразовывает под запись в набор байт
func (e *SrAbsAnSensData) Encode(_ byte) ([]byte, error) {
	return []byte{
		e.SensorNumber,
		byte(e.Value),
		byte(e.Value >> 8),
		byte(e.Value >> 16),
	}, nil
}

// Length получает длину закодированной под записи
func (e *SrAbsAnSensData) Length(_ byte) uint16 {
	return 4
}
