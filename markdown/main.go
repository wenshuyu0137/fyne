package main

import (
	"io/ioutil"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	edit_widget    *widget.Entry
	previer_widget *widget.RichText
	cur_file       fyne.URI
	save_menu_item *fyne.MenuItem
}

func (c *config) make_ui() (edit *widget.Entry, preview *widget.RichText) {
	edit = widget.NewMultiLineEntry()
	preview = widget.NewRichTextFromMarkdown("")

	c.edit_widget = edit
	c.previer_widget = preview

	edit.OnChanged = preview.ParseMarkdown
	return
}

var filter = storage.NewExtensionFileFilter([]string{
	".md",
	".MD",
})

func (c *config) open(win fyne.Window) func() {
	return func() {
		open_dialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {

			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if read == nil {
				return
			}

			defer read.Close()

			data, err := ioutil.ReadAll(read)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			c.edit_widget.SetText(string(data))

			c.cur_file = read.URI()
			win.SetTitle(win.Title() + "_" + read.URI().Name())
			c.save_menu_item.Disabled = false

		}, win)
		open_dialog.SetFilter(filter)

		open_dialog.Show()
	}
}

func (c *config) save(win fyne.Window) func() {
	return func() {
		if c.cur_file != nil {
			write, err := storage.Writer(c.cur_file)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			write.Write([]byte(c.edit_widget.Text))
			defer write.Close()
		}
	}
}

func (c *config) save_as(win fyne.Window) func() {
	return func() {
		save_as_dialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if write == nil {
				//user cancelled
				return
			}

			if !strings.HasSuffix(strings.ToLower(write.URI().String()), ".md") {
				dialog.ShowInformation("Error", "Please name your file with a .md extension!", win)
			}

			//save file
			write.Write([]byte(c.edit_widget.Text))
			c.cur_file = write.URI()

			defer write.Close()

			win.SetTitle(win.Title() + "_" + write.URI().Name())
			c.save_menu_item.Disabled = false
		}, win)
		save_as_dialog.SetFileName("Untitled.md")
		save_as_dialog.SetFilter(filter)
		save_as_dialog.Show()
	}
}

func (c *config) create_menu_items(win fyne.Window) {
	open_menu := fyne.NewMenuItem("Open...", c.open(win))

	save_menu := fyne.NewMenuItem("Save", c.save(win))
	c.save_menu_item = save_menu
	c.save_menu_item.Disabled = true

	save_as_menu := fyne.NewMenuItem("Save as...", c.save_as(win))

	file_menu := fyne.NewMenu("File", open_menu, save_menu, save_as_menu)

	menu := fyne.NewMainMenu(file_menu)

	win.SetMainMenu(menu)
}

var cfg config

func main() {
	//create a fyne app
	a := app.New()

	//create a window for the app
	win := a.NewWindow("MarkDown")

	//get the user interface
	edit, preview := cfg.make_ui()
	cfg.create_menu_items(win)

	//set the content of the window
	win.SetContent(container.NewHSplit(edit, preview))

	//show window and run app
	win.Resize(fyne.Size{
		Width:  800,
		Height: 500,
	})
	win.CenterOnScreen()
	win.ShowAndRun()
}
