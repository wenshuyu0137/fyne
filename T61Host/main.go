package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"os"
	"path/filepath"
)

type main_window struct {
	app_window fyne.Window //应用窗口
}

// 窗口实例化
func new_window() fyne.Window {
	m_app := app.New()

	path := os.Args[0]
	title := filepath.Base(path)
	title_ext := filepath.Ext(path)
	title = title[:len(title)-len(title_ext)]

	m_window := m_app.NewWindow(title)

	m_window.Resize(fyne.NewSize(1000, 600))
	m_window.CenterOnScreen()
	return m_window
}

// 全局变量
var (
	m_app_window = main_window{
		app_window: new_window(),
	}

	m_menu = all_menu{
		all_menu: new_main_menu(),
	}

	serial_cfg = new_config()
)

func main() {
	create_menu()
	m_app_window.app_window.ShowAndRun()
}
