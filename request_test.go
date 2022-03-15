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
	r := DLT2007("130611038998", "B621")
	fmt.Println("DLT2007 frame", hex.EncodeToString(r))
}

func TestCheckSum(t *testing.T) {
	// fmt.Println(hex.EncodeToString(CheckSum([]byte{0x68, 0x89, 0x03, 0x03, 0x11, 0x06, 0x13, 0x68, 0x01, 0x02, 0x44, 0xE9})))
	fmt.Println(CheckSum([]byte{0x68, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x68, 0x01, 0x02, 0x44, 0xE9}))

}

func TestResonseParse(t *testing.T) {
	//      8a 56
	//-33H  57 23
	// b := []byte{0xfe, 0xfe, 0xfe, 0xfe, 0x68, 0x29, 0x25, 0x07, 0x07, 0x21, 0x20, 0x68, 0x91, 0x06, 0x33, 0x34, 0x34, 0x35, 0x8a, 0x56, 0x2e, 0x16}
	// fmt.Println(Response2007(b))
}

func TestDLT1997(t *testing.T) {
	fmt.Println(hex.EncodeToString(DLT1997("130611039061", "B611")))
}
