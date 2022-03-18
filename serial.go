package glt

import (
	"encoding/hex"
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
		Timeout:  2 * time.Second,
	})

	s.Port = port

	if err != nil {
		return nil, err
	}

	return port, nil
}

func (s *SerialPort) Read() ([]byte, error) {
	res := make([]byte, 0)

	for {
		b := make([]byte, 22)
		n, err := DLT645Master.Port.Read(b)
		if err != nil {
			return nil, err
		}
		res = append(res, b[0:n]...)

		if n >= 1 && b[n-1] == byte(22) { //0x16为结束符
			break
		}
	}
	fmt.Println("res:", hex.EncodeToString(res))
	return res, nil
}
