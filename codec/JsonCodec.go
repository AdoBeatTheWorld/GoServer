package codec

import "encoding/json"

type jsonCodec struct {
}

func NewJsonCode() ICodec {
	return &jsonCodec{}
}

func (jc *jsonCodec) Encode(data interface{}) (interface{}, error) {
	buf, err := json.Marshal(data)
	return buf, err
}

func (jc *jsonCodec) Decode(datum interface{}) interface{} {
	return nil
}

func (jc *jsonCodec) GetType() CodecType {
	return CodecJson
}
