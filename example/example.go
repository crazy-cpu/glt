package main

import (
	"fmt"

	glt "github.com/programmer-liu/GLT"
)

func main() {
	// go func() {
	glt.NewSerial(&glt.SerialPort{
		Address:  "/dev/ttyS1",
		BaudRate: 1200,
		DataBit:  8,
		StopBits: 1,
		Parity:   "E",
	})

	//读取A相电流
	no, value := glt.DLT645Master.Request(glt.Protocol1997, "000000038996", "B611")
	fmt.Println("no:", no, " value:", value)

}
