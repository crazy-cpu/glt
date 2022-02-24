package glt

import (
	"encoding/hex"
	"fmt"
	"sync"
)

func (s *SerialPort) Request(address string, dataMarker string) (no string, vlaue float64) {
	s.Locker.Lock()
	defer s.Locker.Unlock()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		DLT645Master.SlaveResponseFrame <- s.Read()
		wg.Done()
	}()

	body := DLT2007(address, dataMarker)
	s.Port.Write(body)
	wg.Wait()

	// DLT645Master.MasterRequestFrame <- data
	res := <-DLT645Master.SlaveResponseFrame
	fmt.Println("res:", res)
	return ResponseParse(res, dataMarker)
}

func DLT2007(address string, dataMarker string) []byte {
	var frame = make([]byte, 0)
	addr, _ := BytesReverse(address)
	marker, _ := BytesReverse(dataMarker)
	frame = append(frame, []byte{0xFE, 0xFE, 0xFE, 0xFE}...) //前置唤醒
	frame = append(frame, 0x68)                              //起始符
	frame = append(frame, addr...)                           //地址域
	frame = append(frame, 0x68)                              //起始符
	frame = append(frame, 0x11)                              //控制码
	frame = append(frame, 0x04)                              //数据长度域
	frame = append(frame, Add33H(marker)...)                 //数据域

	sum := CheckSum(frame[4:18])

	frame = append(frame, sum)  //校验码
	frame = append(frame, 0x16) //结束符
	return frame
}

func BytesReverse(str string) ([]byte, error) {
	if len(str)%2 != 0 {
		return nil, fmt.Errorf("invalid hex str")
	}
	raw, _ := hex.DecodeString(str)
	for i, j := 0, len(raw)-1; i <= j; i, j = i+1, j-1 {
		raw[i], raw[j] = raw[j], raw[i]
	}

	return raw, nil
}

func CheckSum(bytes []byte) byte {
	var sum byte
	for _, b := range bytes {
		sum = sum + (b % 0xFF)
	}
	return sum
}

func Add33H(marker []byte) []byte {
	for i, _ := range marker {
		marker[i] += 0x33
	}

	return marker
}
