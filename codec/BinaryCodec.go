package codec

type binCodec struct {
}

func NewBinCode() ICodec {
	return &binCodec{}
}

func (bc *binCodec) Encode(data interface{}) interface{} {
	return nil
}

func (bc *binCodec) Decode(datum []byte) interface{} {
	return nil
}

func (bc *binCodec) GetType() CodecType {
	return CodecBinary
}
