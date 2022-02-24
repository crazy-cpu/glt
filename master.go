package glt

var DLT645Master *Master

type Master struct {
	MasterRequestFrame []byte

	SlaveResponseFrame chan []byte

	SerialPort
}

func init() {
	if DLT645Master == nil {
		DLT645Master = &Master{
			MasterRequestFrame: make([]byte, 0),
			SlaveResponseFrame: make(chan []byte, 1),
		}
	}
}

func NewSerial(serial *SerialPort) error {
	_, err := serial.Open()
	if err != nil {
		return err
	}
	DLT645Master.SerialPort = *serial
	return nil
}

// func MasterHub() {
// 	for {
// 		select {
// 		case request := <-DLT645Master.MasterRequestFrame:
// 			DLT645Master.Port.Write(request)
// 		case response := <-DLT645Master.SlaveResponseFrame:
// 			fmt.Println("response:", response)
// 		}
// 	}

// }
