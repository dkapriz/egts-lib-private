package egts

// BinaryData интерфейс для работы с бинарными секциями
type BinaryData interface {
	Decode([]byte, byte) error
	Encode(byte) ([]byte, error)
	Length(byte) uint16
}
