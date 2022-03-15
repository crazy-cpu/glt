package glt

import (
	"encoding/hex"
	"fmt"
	"strconv"

	datamarket "github.com/programmer-liu/GLT/datamarker"
)

// type Response struct{}

func Response(protocol string, response []byte, register string) (no string, value float64) {
	// if len(response) != 22 {
	// 	return
	// }需要验证数据完整性，后期加

	var Data []byte

	if protocol == Protocol2007 {
		No, _ := BytesReverse(hex.EncodeToString(response[5:11]))
		Data, _ = BytesReverse(hex.EncodeToString(Sub33H(response[18:20])))
		noHex := hex.EncodeToString(No)
		v, _ := strconv.ParseFloat(hex.EncodeToString(Data), 32)
		return noHex, v / datamarket.DataMarker2007[register]
	} else if protocol == Protocol1997 {
		No, _ := BytesReverse(hex.EncodeToString(response[1:7]))
		Data, _ = BytesReverse(hex.EncodeToString(Sub33H(response[12:14])))
		noHex := hex.EncodeToString(No)
		v, _ := strconv.ParseFloat(hex.EncodeToString(Data), 32)
		return noHex, v / datamarket.DataMarker1997[register]
	} else {
		panic(fmt.Sprintf("invalid protocol:'%s'", protocol))
	}

	// noHex := hex.EncodeToString(No)
	// return noHex, v / datamarket.DataMarker2007[register]

}

func Sub33H(marker []byte) []byte {
	for i, _ := range marker {
		marker[i] -= 0x33
	}

	return marker
}
