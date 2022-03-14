package main

import (
	"fmt"

	glt "github.com/programmer-liu/GLT"
)

func main() {
	// go func() {
	glt.NewSerial(&glt.SerialPort{
		Address:  "/dev/ttyS2",
		BaudRate: 2400,
		DataBit:  8,
		StopBits: 1,
		Parity:   "E",
	})

	//读取A相电流
	no, value := glt.DLT645Master.Request(glt.Protocol2007, "210312036423", "02020100")
	fmt.Println("no:", no, " value:", value)

}
