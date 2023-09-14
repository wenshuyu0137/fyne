package main

import (
	"fmt"
)

func main() {
	var m_cfg = new_config()

	port_names := m_cfg.get_post_names()
	fmt.Println(port_names)
	m_cfg.set_baudrate(115200)
	m_cfg.set_port("COM8")

	port := new_port(&m_cfg)

	send_data := []byte("hellodsawdwasfsafsd world")

	rcv_data := serial_send_and_rcv(port, send_data)
	fmt.Printf("%s\n", string(rcv_data))
}
