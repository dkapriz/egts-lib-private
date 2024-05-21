package egts

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
)

// SrAccelData структура под записи типа EGTS_SR_ACCEL_DATA, показания акселерометра
type SrAccelData struct {
	StructuresAmount uint8             `json:"SA"`
	AbsoluteTime     uint32            `json:"ATM"`
	AccelDataStruct  []AccelDataStruct `json:"ADSs"`
}

type AccelDataStruct struct {
	RelativeTime    uint16 `json:"RTM"`
	XAxisAccelValue int16  `json:"XAAV"`
	YAxisAccelValue int16  `json:"YAAV"`
	ZAxisAccelValue int16  `json:"ZAAV"`
}

// Decode разбирает байты в структуру под записи
func (e *SrAccelData) Decode(content []byte, _ byte) error {
	var (
		err error
	)
	tmpBuf := make([]byte, 4)
	buf := bytes.NewReader(content)

	if e.StructuresAmount, err = buf.ReadByte(); err != nil || e.StructuresAmount == 0 {
		return fmt.Errorf("не удалось получить количество структур: %v", err)
	}

	if _, err = buf.Read(tmpBuf); err != nil {
		return fmt.Errorf("не удалось получить время проведения измерений первой структуры: %v", err)
	}
	e.AbsoluteTime = binary.LittleEndian.Uint32(tmpBuf)

	tmpBuf = make([]byte, 2)
	for i := 0; i < int(e.StructuresAmount); i++ {
		accelData := AccelDataStruct{}
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("не удалось получить приращение ко времени измерения предыдущей записи: %v", err)
		}
		accelData.RelativeTime = binary.LittleEndian.Uint16(tmpBuf)

		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("не удалось получить значение линейного ускорения по оси X: %v", err)
		}
		accelData.XAxisAccelValue = int16(big.NewInt(0).SetBytes(tmpBuf).Int64())

		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("не удалось получить значение линейного ускорения по оси Y: %v", err)
		}
		accelData.YAxisAccelValue = int16(big.NewInt(0).SetBytes(tmpBuf).Int64())

		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("не удалось получить значение линейного ускорения по оси Z: %v", err)
		}
		accelData.ZAxisAccelValue = int16(big.NewInt(0).SetBytes(tmpBuf).Int64())

		e.AccelDataStruct = append(e.AccelDataStruct, accelData)
	}
	return err
}

// Encode преобразовывает под запись в набор байт
func (e *SrAccelData) Encode(_ byte) ([]byte, error) {
	var (
		err    error
		result []byte
	)
	buf := new(bytes.Buffer)

	if err = buf.WriteByte(e.StructuresAmount); err != nil {
		return result, fmt.Errorf("не удалось записать количество структур: %v", err)
	}

	bytesTmpBuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytesTmpBuf, e.AbsoluteTime)
	if _, err = buf.Write(bytesTmpBuf); err != nil {
		return result, fmt.Errorf("не удалось записать время проведения измерений первой структуры: %v", err)
	}

	bytesTmpBuf = make([]byte, 2)
	for i := 0; i < int(e.StructuresAmount); i++ {
		binary.LittleEndian.PutUint16(bytesTmpBuf, e.AccelDataStruct[i].RelativeTime)
		if _, err = buf.Write(bytesTmpBuf); err != nil {
			return result, fmt.Errorf("не удалось записать приращение ко времени измерения предыдущей записи: %v", err)
		}

		xValue := []byte{uint8(e.AccelDataStruct[i].XAxisAccelValue >> 8), uint8(e.AccelDataStruct[i].XAxisAccelValue & 0xff)}
		if _, err = buf.Write(xValue); err != nil {
			return result, fmt.Errorf("не удалось записать значение линейного ускорения по оси X: %v", err)
		}

		yValue := []byte{uint8(e.AccelDataStruct[i].YAxisAccelValue >> 8), uint8(e.AccelDataStruct[i].YAxisAccelValue & 0xff)}
		if _, err = buf.Write(yValue); err != nil {
			return result, fmt.Errorf("не удалось записать значение линейного ускорения по оси Y: %v", err)
		}

		zValue := []byte{uint8(e.AccelDataStruct[i].ZAxisAccelValue >> 8), uint8(e.AccelDataStruct[i].ZAxisAccelValue & 0xff)}
		if _, err = buf.Write(zValue); err != nil {
			return result, fmt.Errorf("не удалось записать значение линейного ускорения по оси Z: %v", err)
		}
	}
	result = buf.Bytes()

	return result, err
}

// Length получает длину закодированной под записи
func (e *SrAccelData) Length(protocolVersion byte) uint16 {
	var result uint16

	if recBytes, err := e.Encode(protocolVersion); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
