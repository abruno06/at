package at

import (
	"go.bug.st/serial"
)

// <CR><LF> sequence use in AT.
const Sep = "\r\n"

// Ctrl+Z code.
const Sub = "\x1A"

// "AT"
const AT = "AT"

type Modem struct {
	DevicePort string
	Speed      int

	port *serial.Port
}

func Open(d *Modem) error {
	mode := &serial.Mode{
		BaudRate: d.Speed,
	}
	port, err := serial.Open(d.DevicePort, mode)
	if err != nil {
		return err
	}
	//for debug purpose, make Serial port accessible to the Modem struct
	d.port = &port
	return nil

}

func Close(d *Modem) error {
	port := *d.port
	return port.Close()
}

func Port(d *Modem) serial.Port {
	return *d.port
}
