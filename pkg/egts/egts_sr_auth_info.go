package egts

import (
	"bytes"
	"fmt"
	"strings"
)

// SrAuthInfo структура под записи типа EGTS_SR_AUTH_INFO, которая предназначена для передачи на
// телематическую платформу аутентификационных данных АС с использованием ранее переданных
// со стороны платформы параметров для осуществления шифрования данных.
type SrAuthInfo struct {
	UserName       string `json:"UNM"`
	UserPassword   string `json:"UPSW"`
	ServerSequence string `json:"SS"`
}

// Decode разбирает байты в структуру под записи
func (e *SrAuthInfo) Decode(content []byte, _ byte) error {
	var (
		err    error
		tmpStr string
	)
	//разделитель строковых полей из ГОСТ 54619 - 2011 секции EGTS_SR_AUTH_INFO
	sep := byte(0x00)

	buf := bytes.NewBuffer(content)
	tmpStr, err = buf.ReadString(sep)
	if err != nil {
		return fmt.Errorf("не удалось считать имя пользователя sr_auth_info: %v", err)
	}
	e.UserName = strings.TrimSuffix(tmpStr, string(sep))

	tmpStr, err = buf.ReadString(sep)
	if err != nil {
		return fmt.Errorf("не удалось считать пароль sr_auth_info: %v", err)
	}
	e.UserPassword = strings.TrimSuffix(tmpStr, string(sep))

	if buf.Len() > 0 {
		tmpStr, err = buf.ReadString(sep)
		if err != nil {
			return fmt.Errorf("не удалось считать SS из sr_auth_info: %v", err)
		}
		e.ServerSequence = strings.TrimSuffix(tmpStr, string(sep))
	}

	return err
}

// Encode преобразовывает под запись в набор байт
func (e *SrAuthInfo) Encode(_ byte) ([]byte, error) {
	var (
		err    error
		result []byte
	)
	//разделитель строковых полей из ГОСТ 54619 - 2011 секции EGTS_SR_AUTH_INFO
	sep := byte(0x00)

	result = append(result, []byte(e.UserName)...)
	result = append(result, sep)

	result = append(result, []byte(e.UserPassword)...)
	result = append(result, sep)

	// необязательное поле, наличие зависит от используемого алгоритма шифрования
	// специальная серверная последовательность байт, передаваемая в теле под записи EGTS_SR_AUTH_PARAMS
	if e.ServerSequence != "" {
		result = append(result, []byte(e.ServerSequence)...)
		result = append(result, sep)
	}

	return result, err
}

// Length получает длину закодированной под записи
func (e *SrAuthInfo) Length(protocolVersion byte) uint16 {
	var result uint16

	if recBytes, err := e.Encode(protocolVersion); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
