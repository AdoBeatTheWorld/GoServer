package codec

import (
	"fmt"
	"github.com/pkg/errors"
	"sync/atomic"
)

type ICodec interface {
	Encode(data interface{}) (interface{}, error)
	Decode(datum interface{}) interface{}
	GetType() CodecType
}

type CodecType uint8

type codecManager struct {
	//sync.Mutex
	codecMap map[CodecType]ICodec
}

var manager codecManager

const (
	CodecText CodecType = iota
	CodecJson
	CodecBinary
	CodecProto
	CodecSproto
)

var flag uint32 = 0

func InitCodec() error {
	if atomic.LoadUint32(&flag) == 1 {
		return errors.New("Codec already initialized...")
	}

	manager = codecManager{
		codecMap: make(map[CodecType]ICodec),
	}
	atomic.StoreUint32(&flag, 1)
	return nil
}

func RegisterCodec(ctype CodecType, codec ICodec) error {
	if manager.codecMap[ctype] != nil {
		return errors.New(fmt.Sprintf("Illegal Codec Type:%d", ctype))
	}
	manager.codecMap[ctype] = codec
	return nil
}

func GetCodec(ctype CodecType) (ICodec, error) {
	if manager.codecMap[ctype] != nil {
		return manager.codecMap[ctype], nil
	}
	return nil, errors.New(fmt.Sprintf("Illegal CodecType:%d", ctype))
}
