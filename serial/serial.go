package serial

import "go.bug.st/serial"

// opens port - opens port at 9600 bps, returns the port and err
func OpenPort(baud int, portName string) (serial.Port, error) {
	mode := &serial.Mode{BaudRate: baud}
	port, err := serial.Open(portName, mode)
	return port, err
}

// sends command - writes a string to the port
func SendCommand(port serial.Port, command string) error {
	_, err := port.Write([]byte(command + "\n"))
	return err
}
