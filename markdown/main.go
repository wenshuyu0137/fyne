package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)


type config struct {
	edit_widget *widget.Entry
	previer_widget *widget.RichText
	cur_file fyne.URI
	save_menu_item *fyne.MenuItem
}

func (c *config)make_ui() (edit *widget.Entry,preview *widget.RichText){
	edit = widget.NewMultiLineEntry()
	preview = widget.NewRichTextFromMarkdown("")

	c.edit_widget = edit
	c.previer_widget = preview

	edit.OnChanged = preview.ParseMarkdown
	return
}

func (c *config)create_menu_window(win fyne.Window){
	open_menu := fyne.NewMenuItem("Open...",func(){

	})

	save_menu := fyne.NewMenuItem("Save",func(){

	})

	save_as_menu := fyne.NewMenuItem("Save as...",func(){

	})

	file_menu := fyne.NewMenu("File",open_menu,save_menu,save_as_menu)

	menu := fyne.NewMainMenu(file_menu)

	win.SetMainMenu(menu)
}

var cfg config

func main(){
	//create a fyne app
	a := app.New()	

	//create a window for the app
	win := a.NewWindow("MarkDown")

	//get the user interface
	edit,preview := cfg.make_ui()
	cfg.create_menu_window(win)

	//set the content of the window
	win.SetContent(container.NewHSplit(edit,preview))

	//show window and run app
	win.Resize(fyne.Size{
		Width:800,
		Height:500,
	})
	win.CenterOnScreen()
	win.ShowAndRun()
}

