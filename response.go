package glt

import (
	"encoding/hex"
	"fmt"
	"strconv"

	datamarker "github.com/programmer-liu/GLT/datamarker"
)

// type Response struct{}

func Response(protocol string, response []byte, register string) (no string, value float64, err error) {
	if !VerifyResponse(response) {
		return "", 0, fmt.Errorf("response content incorrect of slave machine--->%s",
			hex.EncodeToString(response))
	}

	res := IgnoreFrontGuide(response)
	number, _ := BytesReverse(hex.EncodeToString(res[1:7]))
	actualValue := GetValue(protocol, res)
	data, _ := BytesReverse(hex.EncodeToString(actualValue))
	v, _ := strconv.ParseFloat(hex.EncodeToString(data), 32)

	decimal, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", v), 32)

	if protocol == Protocol1997 {
		return hex.EncodeToString(number), decimal / datamarker.DataMarker1997[register], nil
	} else if protocol == Protocol2007 {
		return hex.EncodeToString(number), decimal / datamarker.DataMarker2007[register], nil
	}

	return "", 0, fmt.Errorf("response abnormal!")
}

/*
	DL/T-1997和DL/T-2007第9个字节均代表实际参数长度

*/
func GetValue(protocol string, b []byte) []byte {
	l := b[9]
	if protocol == Protocol1997 {
		return Sub33H(b[9+2+1 : 9+l+1])
	} else if protocol == Protocol2007 {
		if l < 4 {
			return Sub33H(b[9+1 : 9+l+1])
		}
		return Sub33H(b[9+4+1 : 9+l+1])
	} else {
		panic(fmt.Sprintf("invalid protocol type:%s", protocol))
	}
}

func IgnoreFrontGuide(b []byte) []byte {
	if len(b) == 0 || b[0] != 0xFE {
		return b
	}

	var cnt int = 1
	length := len(b)
	//首字节必须为0xFE,最多4个连续0xFE
	for i := 1; i < 4; i++ {
		if b[i] != 0xFE {
			break
		}

		cnt++
	}
	return b[cnt:length]
}

func VerifyResponse(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	res := IgnoreFrontGuide(b)
	if res[0] != 0x68 || res[len(res)-1] != 0x16 || CheckSum(res[:len(res)-2]) != res[len(res)-2] {
		return false
	}
	return true
}

func Sub33H(marker []byte) []byte {
	for i, _ := range marker {
		marker[i] -= 0x33
	}

	return marker
}
