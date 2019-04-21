package test

import (
	"github.com/golang/protobuf/proto"
	"gitlab.com/adoontheway/goserver/test/proto"
	"testing"
)

func TestProto(t *testing.T) {
	h := &hello.Hello{Content: "it is ok"}
	data, err := proto.Marshal(h)
	if err != nil {
		t.FailNow()
	}
	h1 := &hello.Hello{}
	err = proto.Unmarshal(data, h1)
	if err != nil {
		t.FailNow()
	}

	if h.Content != h1.Content {
		t.FailNow()
	}
}
