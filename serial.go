package glt

import (
	"fmt"
	"sync"
	"time"

	"github.com/goburrow/serial"
)

type SerialPort struct {
	Locker   sync.Mutex
	Address  string
	Port     serial.Port
	BaudRate int
	DataBit  int
	StopBits int
	Parity   string
}

func (s *SerialPort) Open() (serial.Port, error) {
	port, err := serial.Open(&serial.Config{
		Address:  s.Address,
		BaudRate: s.BaudRate,
		DataBits: s.DataBit,
		StopBits: s.StopBits,
		Parity:   s.Parity,
		Timeout:  500 * time.Millisecond,
	})

	s.Port = port

	if err != nil {
		return nil, err
	}

	return port, nil
}

func (s *SerialPort) Read() []byte {
	res := make([]byte, 0)

	for {

		b := make([]byte, 22)
		n, err := DLT645Master.Port.Read(b)
		if err != nil {
			fmt.Println("err:", err)
			return nil
		}
		res = append(res, b[0:n]...)
		fmt.Println("res:", res, "n :", n)

		// if len(res)==16{
		// 	break
		// }
		if b[n-1] == byte(22) { //0x16为结束符
			break
		}
	}

	return res
}
