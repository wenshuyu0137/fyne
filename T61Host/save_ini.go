package main

import (
	"errors"
	"github.com/go-ini/ini"
	"os"
)

var ini_path = "./config.ini"

// 保存配置
func save_to_init(section string, option string, value string) {
	if _, err := os.Stat(ini_path); os.IsNotExist(err) {
		cfg := ini.Empty()
		err := cfg.SaveTo(ini_path)
		if err != nil {
			return
		}
	}
	cfg, _ := ini.Load(ini_path)

	sec := cfg.Section(section)
	op := sec.Key(option)
	op.SetValue(value)
	_ = cfg.SaveTo(ini_path)
}

// 加载配置
func load_from_ini(section string, option string) (string, error) {
	if _, err := os.Stat(ini_path); os.IsNotExist(err) {
		return "", nil
	}
	cfg, _ := ini.Load(ini_path)

	sec, err := cfg.GetSection(section)
	if err != nil {
		return "", nil
	}

	value := sec.Key(option).String()
	if value == "" {
		return "", errors.New("no option")
	}

	return value, nil
}
