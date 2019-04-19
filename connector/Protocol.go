package connector

type Header struct {
	PackLen   uint16
	CRC       uint16
	Version   uint16
	Sign      uint16
	MainId    uint8
	SubId     uint8
	EncryType uint8
	Back      uint8
	RequestId uint32
	DataSize  uint16
	Data      []byte
}
