package glt

import (
	"fmt"
)

func RequestFrameAssemble2007(address string, dataMarker string) []byte {
	var frame []byte
	addr, _ := BytesReverse(address)
	marker, _ := BytesReverse(dataMarker)
	fmt.Println("addr:", addr, " market:", marker)
	frame = append(frame, []byte{0xFE, 0xFE, 0xFE, 0xFE}...) //前置唤醒
	frame = append(frame, 0x68)                              //起始符
	frame = append(frame, addr...)                           //地址域
	frame = append(frame, 0x68)                              //起始符
	frame = append(frame, 0x11)                              //控制码
	frame = append(frame, 0x04)                              //数据长度域
	frame = append(frame, marker...)                         //数据域
	sum := CheckSum(frame[:18])

	frame = append(frame, sum)  //校验码
	frame = append(frame, 0x16) //结束符

	return frame
}

func BytesReverse(str string) ([]byte, error) {
	if len(str)%2 != 0 {
		return nil, fmt.Errorf("invalid hex str")
	}

	raw := []rune(str)

	for i, j := 0, len(str)-2; i < j; i, j = i+2, j-2 {
		raw[i], raw[i+1], raw[j], raw[j+1] = raw[j], raw[j+1], raw[i], raw[i+1]
	}
	return []byte(string(raw)), nil
}

func CheckSum(bytes []byte) byte {
	var sum byte
	for _, b := range bytes {
		sum = sum + (b % 0xFF)
	}
	return sum
}
