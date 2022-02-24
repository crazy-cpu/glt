package glt

import (
	"fmt"
	"sync"

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
	})

	s.Port = port

	if err != nil {
		return nil, err
	}

	return port, nil
}

func (s *SerialPort) Read() []byte {
	res := make([]byte, 0)
	// var Bytes = 0
	for {

		b := make([]byte, 22)
		n, _ := DLT645Master.Port.Read(b)
		// Bytes = Bytes + n
		res = append(res, b[0:n]...)
		fmt.Println("Res:", res, " finishFlag:", b[n-1])
		if b[n-1] == byte(22) { //0x16为结束符
			break
		}
	}

	// DLT645Master.SlaveResponseFrame <- b
	return res
}
