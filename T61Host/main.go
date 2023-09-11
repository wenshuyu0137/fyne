package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	var m_cfg = new_config()

	port_names := m_cfg.get_post_names()

	m_cfg.set_baudrate(115200)
	m_cfg.set_port(port_names[0])

	fmt.Printf("baud:%d\n", m_cfg.mode.BaudRate)
	fmt.Printf("port:%v\n", m_cfg.port_name)

	port := new_port(&m_cfg)
	fmt.Printf("%v\n", port)

	len_wr, err := port.Write([]byte("hello world"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len_wr)

	port.Drain()

	rd_buf := make([]byte, 128)
	for {
		n, err := port.Read(rd_buf)
		if err != nil {
			fmt.Println(err)
		}
		if n == 0{
			fmt.Println("EOF")
			break
		}
		fmt.Printf("%v",string(rd_buf[:n]))
	}

}















