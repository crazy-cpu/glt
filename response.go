package glt

import (
	datamarket "GLT/datamarker"
	"encoding/hex"
	"fmt"
	"strconv"
)

type Response struct{}

func ResponseParse(response []byte,register string) (no string, value float64) {
	// if len(response) != 22 {
	// 	return
	// }需要验证数据完整性，后期加

	No, _ := BytesReverse(hex.EncodeToString(response[5:11]))
	Data, _ := BytesReverse(hex.EncodeToString(Sub33H(response[18:20])))
	
	v, _ := strconv.ParseFloat(hex.EncodeToString(Data), 32)
	noHex := hex.EncodeToString(No)
	fmt.Println("v:", v, " dataMarker:", datamarket.DataMarker[register], "result:", v/datamarket.DataMarker[register])
	return noHex, v / datamarket.DataMarker[register]

}

func Sub33H(marker []byte) []byte {
	for i, _ := range marker {
		marker[i] -= 0x33
	}

	return marker
}
