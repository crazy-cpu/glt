package glt

import (
	"fmt"
	"testing"
)

func TestBytesReverse(t *testing.T) {
	str := "1A2B3C4D5E6F"
	res, err := BytesReverse(str)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(string(res))
}

func TestCheckSum(t *testing.T) {
	fmt.Println(CheckSum([]byte{0x68,0x23,0x64,0x03,0x12,0x03,0x21,0x68,0x11,0x04,0x33,0x34,0x34,0x35}))
	r := RequestFrameAssemble2007("210312036423", "02010100")
	fmt.Println(r)
}


