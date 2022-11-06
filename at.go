package at

import (
	"log"

	"go.bug.st/serial"
)

type Device struct {
	Port  string
	Speed int
}

func Open(d *Device) (serial.Port, error) {
	mode := &serial.Mode{
		BaudRate: d.Speed,
	}
	port, err := serial.Open(d.Port, mode)
	if err != nil {
		log.Fatal(err)
	}
	return port, err
}
