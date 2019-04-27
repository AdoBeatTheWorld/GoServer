package test

import (
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
)

func TestUuid(t *testing.T)  {
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDV4:%s \n",u1)

	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something is wrong:%s \n",err)
		t.Fail()
		return
	}
	fmt.Printf("UUIDV4: %s \n", u2)

	u2,err = uuid.FromString("65e77ded-f19c-46a1-8852-1fdc4a02733f")
	if err != nil {
		fmt.Printf("Something is wrong: %s",err)
		t.Fail()
		return
	}
	fmt.Printf("Successfully parsed:%s",u2)
}
