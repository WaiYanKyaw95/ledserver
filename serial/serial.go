package serial

import (
	"time"

	"go.bug.st/serial"
)

// opens port - opens port at 9600 bps, returns the port and err
func OpenPort(baud int, portName string) (serial.Port, error) {
	mode := &serial.Mode{BaudRate: baud}
	port, err := serial.Open(portName, mode)
	if err != nil {
		return nil, err
	}
	// wait for two seconds after opening the port
	time.Sleep(2 * time.Second)
	return port, nil
}

// sends command - writes a string to the port
func SendCommand(port serial.Port, command string) error {
	_, err := port.Write([]byte(command + "\n"))
	return err
}
