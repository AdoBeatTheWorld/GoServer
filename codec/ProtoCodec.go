package codec

type protoCodec struct {
}

func NewProtoCode() ICodec {
	return &protoCodec{}
}

func (pc *protoCodec) Encode(data interface{}) interface{} {
	return nil
}

func (pc *protoCodec) Decode(datum []byte) interface{} {
	return nil
}

func (pc *protoCodec) GetType() CodecType {
	return CodecProto
}
