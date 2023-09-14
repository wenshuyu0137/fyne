package main

import (
	"fmt"
)

func main() {
	var m_cfg = new_config()

	port_names := m_cfg.get_post_names()

	m_cfg.set_baudrate(115200)
	m_cfg.set_port(port_names[0])

	port := new_port(&m_cfg)

	send_data := []byte("hello world")
	rcv_data := serial_send_and_rcv(port,send_data)
	fmt.Println(rcv_data)
}















