package main

import (
	"fmt"
	"time"

	"go.bug.st/serial"
)

type serial_config struct {
	port_name string
	mode      serial.Mode
}

func (s *serial_config) get_post_names() []string {

	ports, err := serial.GetPortsList()
	if err != nil {
		return []string{}
	}

	return ports

}

func (s *serial_config) set_baudrate(value int) {
	s.mode.BaudRate = value
}

func (s *serial_config) set_port(port string) {
	s.port_name = port
}

func new_config() serial_config {
	return serial_config{
		mode: serial.Mode{
			DataBits: 8,
			Parity:   serial.NoParity,
			StopBits: serial.OneStopBit,
		},
	}
}

func new_port(cfg *serial_config) serial.Port {
	port, err := serial.Open(cfg.port_name, &cfg.mode)
	if err != nil {
		return nil
	}
	err = port.SetReadTimeout(time.Millisecond)
	if err != nil {
		fmt.Println(err)
	}
	return port
}

func ctor_serial(cfg *serial_config) serial.Port {
	port, err := serial.Open(cfg.port_name, &cfg.mode)
	if err != nil {
		return nil
	}
	return port
}
