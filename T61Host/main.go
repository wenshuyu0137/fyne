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

	len_wr, err := port.Write([]byte("hello world"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("write length:",len_wr)

	port.Drain()

	rd_buf := make([]byte, 128)
	for {
		n, err := port.Read(rd_buf)
		if err != nil {
			fmt.Println(err)
		}
		if n == 0{
			break
		}
		fmt.Printf("%v",string(rd_buf[:n]))
	}

}















