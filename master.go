package glt

import "github.com/goburrow/serial"

var DLT645Master *Master

type Master struct {
	// MasterRequestFrame chan []RequestFrame2007

	SlaveResponseFrame chan []byte

	Serial
}

func NewMaster(serial *Serial) {
	if DLT645Master == nil {
		serial.Open()
		DLT645Master = &Master{
			// MasterRequestFrame: make(chan []RequestFrame2007, 50),
			SlaveResponseFrame: make(chan []byte, 50),
			Serial:             *serial,
		}
	}
}

type Serial struct {
	address  string
	Port     serial.Port
	BaudRate int
	DataBit  int
	StopBits int
	Parity   string
}

func (s *Serial) Open() (serial.Port, error) {
	port, err := serial.Open(&serial.Config{
		Address:  s.address,
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
