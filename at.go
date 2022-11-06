package at

import (
	"log"

	"go.bug.st/serial"
)

// <CR><LF> sequence use in AT.
const Sep = "\r\n"

// Ctrl+Z code.
const Sub = "\x1A"

// "AT"
const AT = "AT"

type Device struct {
	DevicePort string
	Speed      int

	port *serial.Port
}

func Open(d *Device) error {
	mode := &serial.Mode{
		BaudRate: d.Speed,
	}
	port, err := serial.Open(d.DevicePort, mode)
	if err != nil {
		log.Fatal(err)
	}
	d.port = &port
	return err
}
