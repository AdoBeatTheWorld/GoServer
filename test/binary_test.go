package test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	buf := new(bytes.Buffer)
	leng := buf.Len()
	a := 0
	var b int32 = 0
	c := 1
	d := false
	e := true
	binary.Write(buf, binary.LittleEndian, a)
	fmt.Printf("Len:%d, Cap:%d  [a := 0]\n", buf.Len()-leng, buf.Cap())
	leng = buf.Len()
	binary.Write(buf, binary.LittleEndian, b)
	fmt.Printf("Len:%d, Cap:%d  [var b int32 = 0]\n", buf.Len()-leng, buf.Cap())
	leng = buf.Len()
	binary.Write(buf, binary.LittleEndian, c)
	fmt.Printf("Len:%d, Cap:%d  [c := 1]\n", buf.Len()-leng, buf.Cap())
	leng = buf.Len()
	binary.Write(buf, binary.LittleEndian, d)
	fmt.Printf("Len:%d, Cap:%d  [d := false]\n", buf.Len()-leng, buf.Cap())
	leng = buf.Len()
	binary.Write(buf, binary.LittleEndian, e)
	fmt.Printf("Len:%d, Cap:%d  [e := true]\n", buf.Len()-leng, buf.Cap())

	fmt.Printf("%x \n", buf.Bytes())

	var f int32
	var g, h bool
	binary.Read(buf, binary.LittleEndian, &f)
	binary.Read(buf, binary.LittleEndian, &g)
	binary.Read(buf, binary.LittleEndian, &h)
	fmt.Printf("%d,%v,%v", f, g, h)
}
