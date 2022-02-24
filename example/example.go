package main

import (
	glt "GLT"
	"fmt"
)

func main() {
	// go func() {
	glt.NewSerial(&glt.SerialPort{
		Address:  "/dev/ttyS2",
		BaudRate: 2400,
		DataBit:  8,
		StopBits: 2,
		Parity:   "N",
	})

	//读取A相电流
	no, value := glt.DLT645Master.Request("210312036423", "02020100")
	fmt.Println("no:", no, " value:", value)

}
