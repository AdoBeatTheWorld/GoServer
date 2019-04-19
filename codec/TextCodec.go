package codec

type textCodec struct {
}

func NewTextCodec() ICodec {
	return &textCodec{}
}

func (tc *textCodec) Encode(data interface{}) interface{} {
	return nil
}

func (tc *textCodec) Decode(datum []byte) interface{} {
	return nil
}

func (tc *textCodec) GetType() CodecType {
	return CodecText
}
