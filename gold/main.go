package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type app_config struct {
	app         fyne.App
	info_log    *log.Logger
	err_log     *log.Logger
	main_window fyne.Window
}

var my_app app_config

func main() {
	//create a fyne application
	fyne_app := app.NewWithID("my frist project")
	my_app.app = fyne_app

	my_app.info_log = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	my_app.err_log = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//create our loggers

	//open a connection to the database

	//create a database repository

	//create and size a fyne window
	my_app.main_window = fyne_app.NewWindow("GoldWatcher")
	my_app.main_window.Resize(fyne.NewSize(300, 200))
	my_app.main_window.SetFixedSize(true)
	my_app.main_window.SetMaster()

	my_app.make_ui()

	//show and run the application
	my_app.main_window.ShowAndRun()
}
