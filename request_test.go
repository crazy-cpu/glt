package glt

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestBytesReverse(t *testing.T) {
	str := "02010100"
	res, err := BytesReverse(str)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(res)
}

func TestDLT2007(t *testing.T) {
	r := DLT2007("210312036423", "02010100")
	fmt.Println("DLT2007 frame", hex.EncodeToString(r))
}

func TestCheckSum(t *testing.T){
	
}
