package glt

import (
	"encoding/hex"
	"testing"
)

type Case struct {
	Input    interface{}
	Expected interface{}
	Attach   interface{}
	Descript string
}

func TestIgnoreFrontGuide(t *testing.T) {
	TestCase := []Case{
		{
			Input:    []byte{0xFE, 0xFE, 0xFE, 0xFE, 0x86, 0x11, 0x2c, 0x86, 0x16},
			Expected: []byte{0x86, 0x11, 0x2c, 0x86, 0x16},
			Descript: "连续4个前置唤醒字节",
		},
		{
			Input:    []byte{},
			Expected: []byte{},
			Descript: "空响应",
		},
		{
			Input:    []byte{0xFE, 0x86, 0x11, 0x2c, 0x86, 0x16},
			Expected: []byte{0x86, 0x11, 0x2c, 0x86, 0x16},
			Descript: "一个前置唤醒字节",
		},
		{
			Input:    []byte{0xFE, 0x86, 0xFE, 0x2c, 0x86, 0x16},
			Expected: []byte{0x86, 0xFE, 0x2c, 0x86, 0x16},
			Descript: "不连续的前置唤醒字节",
		},
		{
			Input:    []byte{0x86, 0x11, 0x2c, 0x86, 0x16},
			Expected: []byte{0x86, 0x11, 0x2c, 0x86, 0x16},
			Descript: "无前置唤醒字节",
		},
	}

	for _, v := range TestCase {
		actual := IgnoreFrontGuide(v.Input.([]byte))
		if !EqualSlice(actual, v.Expected.([]byte)) {
			t.Errorf("descript:%s input:%s expected:%s actual:%s",
				v.Descript,
				hex.EncodeToString(v.Input.([]byte)),
				hex.EncodeToString(v.Expected.([]byte)),
				hex.EncodeToString(actual))
		}
	}
}

func TestVerifyResponse(t *testing.T) {
	TestCase := []Case{
		{
			Input:    []byte{0x68, 0x61, 0x01, 0x00, 0x00, 0x00, 0x00, 0x68, 0x91, 0x06, 0x33, 0x34, 0x34, 0x35, 0x34, 0x43, 0x10, 0x16},
			Expected: true,
			Descript: "校验和正确的响应数据",
		},
		{
			Input:    []byte{0x68, 0x61, 0x01, 0x00, 0x00, 0x00, 0x00, 0x68, 0x91, 0x06, 0x33, 0x34, 0x34, 0x35, 0x34, 0x43, 0x20, 0x16},
			Expected: false,
			Descript: "校验和错误的响应数据",
		},
		{
			Input:    []byte{0x68, 0x61, 0x01, 0x00, 0x00, 0x00, 0x00, 0x68, 0x91, 0x06, 0x33, 0x34, 0x34, 0x35, 0x34, 0x43, 0x10},
			Expected: false,
			Descript: "缺少结束符0x16",
		},
	}

	for _, v := range TestCase {
		result := VerifyResponse(v.Input.([]byte))
		if v.Expected != VerifyResponse(v.Input.([]byte)) {
			t.Errorf("descript:%s input:%s expected:%v actual:%v",
				v.Descript,
				hex.EncodeToString(v.Input.([]byte)),
				v.Expected.(bool),
				result)
		}
	}
}

func TestGetValue(t *testing.T) {
	TestCase := []Case{
		{
			Input: []byte{0x68, 0x23, 0x64, 0x03, 0x12, 0x03, 0x21, 0x68,
				0x91, 0x07, 0x33, 0x34, 0x35, 0x35, 0x97, 0x34, 0x33, 0xf7, 0x16},
			//9+4+1 : 9+l+1 ->    14:17
			Expected: []byte{0x64, 0x01, 0x00},
			Attach:   "DL/T-2007",
			Descript: "DL/T-2007获取A相电流，占用3字节，减去33H",
		},
		{
			Input: []byte{0x68, 0x61, 0x01, 0x00, 0x00, 0x00, 0x00, 0x68,
				0x91, 0x06, 0x33, 0x34, 0x34, 0x35, 0x34, 0x43, 0x10, 0x16},
			//9+4+1 : 9+l+1 ->    14:17
			Expected: []byte{0x01, 0x10},
			Attach:   "DL/T-2007",
			Descript: "DL/T-2007获取A相电压，占用2字节，减去33H",
		},

		{
			Input: []byte{0x68, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x68,
				0x81, 0x04, 0x54, 0xE9, 0xCC, 0x37, 0x98, 0x16},
			//9+l+1-2 : 9+l+1     12:14
			Expected: []byte{0x99, 0x04},
			Attach:   "DL/T-1997",
			Descript: "DL/T-1997获取A相电流，占用2字节，减去33H",
		},
		{
			Input: []byte{0x68, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x68,
				0x81, 0x05, 0x11, 0xB6, 0xCC, 0x37, 0x98, 0x16},
			//9+l+1-2 : 9+l+1     12:14
			Expected: []byte{0x99, 0x04, 0x65},
			Attach:   "DL/T-1997",
			Descript: "DL/T-1997获取A相电压，占用3字节，减去33H",
		},
	}

	for _, v := range TestCase {
		act := GetValue(v.Attach.(string), v.Input.([]byte))
		if !EqualSlice(act, v.Expected.([]byte)) {
			t.Errorf("descript:%s input:%v expected:%v actual:%v",
				v.Descript,
				hex.EncodeToString(v.Input.([]byte)),
				v.Expected.([]byte),
				act)
		}
	}
}
func EqualSlice(s1, s2 []byte) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
