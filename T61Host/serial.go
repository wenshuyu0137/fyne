package main

import (
	"errors"
	"fmt"
	"time"

	"go.bug.st/serial"
)

// 串口相关结构体
type serial_config struct {
	port_name string      //当前串口号
	mode      serial.Mode //串口的配置
	port      serial.Port //串口的接口类型
	is_open   bool
}

// 获取所有的串口号
func (s *serial_config) get_post_names() []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		return []string{}
	}
	return ports
}

// 设置波特率
func (s *serial_config) set_baudrate(value int) {
	s.mode.BaudRate = value
}

// 设置串口号
func (s *serial_config) set_port(port string) {
	s.port_name = port
}

// 串口配置实例化
func new_config() *serial_config {
	return &serial_config{
		mode: serial.Mode{
			DataBits: 8,
			Parity:   serial.NoParity,
			StopBits: serial.OneStopBit,
		},
	}
}

// 连接串口,接口实例化
func (s *serial_config) connect_port() error {
	port, err := serial.Open(s.port_name, &s.mode)
	if err != nil {
		s.is_open = false
		return err
	}
	s.is_open = true
	s.port = port
	s.is_open = true
	err = port.SetReadTimeout(time.Millisecond)
	if err != nil {
		return err
	}
	return nil
}

func (s *serial_config) close_port() error {
	if s.port != nil {
		err := s.port.Close()
		if err != nil {
			return errors.New("close fail")
		}
		s.is_open = false
	}
	return nil
}

// 串口连接的槽函数
func connect_serial_com(com string, baud int) bool {
	if serial_cfg.port != nil {
		_ = serial_cfg.port.Close()
	}
	serial_cfg.set_port(com)
	serial_cfg.set_baudrate(baud)
	err := serial_cfg.connect_port()
	if err != nil {
		fmt.Println("串口被占用")
		return false
	}
	return true
}

// 串口的发送并接收函数
func serial_send_and_rcv(port serial.Port, data []byte) ([]byte, error) {
	length, err := port.Write(data)
	if err != nil && length != len(data) {
		return nil, err
	}

	err = port.Drain()
	if err != nil {
		return nil, err
	}
	rcv_buf := make([]byte, 256)
	rcv_data := make([]byte, 0)
	for {
		n, err := port.Read(rcv_buf)
		if err != nil {
			return nil, err
		}
		if n == 0 {
			break
		}
		rcv_data = append(rcv_data, rcv_buf[:n]...)
	}
	return rcv_data, nil
}
