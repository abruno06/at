package at

import (
	"log"

	"go.bug.st/serial"
)

func Open() serial.Port {
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		log.Fatal(err)
	}
	return port
}
