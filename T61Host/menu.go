package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	dialog2 "fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

// 主菜单
type all_menu struct {
	all_menu *fyne.MainMenu
}

// 添加子菜单的方法
func (a *all_menu) add_new_menu(menu *fyne.Menu) {
	a.all_menu.Items = append(a.all_menu.Items, menu)
}

// 构造函数,构造主菜单
func new_main_menu() *fyne.MainMenu {
	menu := fyne.NewMainMenu()
	return menu
}

// 弹出串口连接的窗口
func show_serial_dialog(com_name []string, baud []string) {
	var dialog *dialog2.CustomDialog

	com_label := widget.NewLabel("COM")
	com_dropdown := widget.NewSelect(com_name, nil)
	ini_com, err := load_from_ini("串口信息", "COM口")
	if err == nil {
		com_dropdown.SetSelected(ini_com)
	}

	hbox1 := container.NewGridWithRows(1, com_label, com_dropdown)

	baud_label := widget.NewLabel("baudrate")
	baud_dropdown := widget.NewSelect(baud, nil)
	ini_baud, err := load_from_ini("串口信息", "波特率")
	if err == nil {
		baud_dropdown.SetSelected(ini_baud)
	}
	hbox2 := container.NewGridWithRows(1, baud_label, baud_dropdown)

	connect_btn := widget.NewButton("Connect", func() {

		cur_com := com_dropdown.Selected
		cur_baud := baud_dropdown.Selected
		baud_value, _ := strconv.Atoi(cur_baud)

		ret := connect_serial_com(cur_com, baud_value)
		if ret {
			//保存配置文件
			save_to_init("串口信息", "COM口", cur_com)
			save_to_init("串口信息", "波特率", cur_baud)

			dialog.Hide()
		}

	})
	disconnect_btn := widget.NewButton("Disconnect", func() {
		err := serial_cfg.close_port()
		if err != nil {
			return
		} else {
			dialog.Hide()
		}
	})
	hbox3 := container.NewGridWithRows(3, connect_btn, disconnect_btn)

	content := container.NewVBox()
	content.Add(hbox1)
	content.Add(hbox2)
	content.Add(hbox3)

	dialog = dialog2.NewCustom("Serial Connect", "Cancel", content, m_app_window.app_window)
	dialog.Show()
}

// 生成主菜单,并添加相关的子菜单
func create_menu() {

	//--------------------------------串口菜单
	com := serial_cfg.get_post_names()
	baud := []string{"115200", "9600"}
	connect_item := fyne.NewMenuItem("connect", func() {
		show_serial_dialog(com, baud)
	})
	serial_menu := fyne.NewMenu("Serial", connect_item)
	//--------------------------------

	m_menu.add_new_menu(serial_menu)

	m_app_window.app_window.SetMainMenu(m_menu.all_menu)
}
