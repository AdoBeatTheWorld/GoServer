package codec

import "github.com/davyxu/protoplus/proto"

type protoCodec struct {
}

func NewProtoCode() ICodec {
	return &protoCodec{}
}

func (pc *protoCodec) Encode(data interface{}) (interface{}, error) {
	buf, err := proto.Marshal(data)
	return buf, err
}

func (pc *protoCodec) Decode(datum interface{}) interface{} {
	return nil
}

func (pc *protoCodec) GetType() CodecType {
	return CodecProto
}
