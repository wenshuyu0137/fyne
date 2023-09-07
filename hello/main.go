package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type test_app struct{
	ouput *widget.Label
}

var my_app test_app


func (t *test_app) make_ui() (*widget.Label,*widget.Entry,*widget.Button){
	ouput := widget.NewLabel("hello world")
	entry := widget.NewEntry()
	btn := widget.NewButton("Entry",func() {
		t.ouput.SetText(entry.Text)
	})
	btn.Importance = widget.HighImportance

	t.ouput = ouput
	
	return ouput,entry,btn
}

func main(){
	a := app.New()
	w := a.NewWindow("hello world!")

	ouput,entry,btn := my_app.make_ui()

	w.SetContent(container.NewVBox(ouput,entry,btn))
	w.Resize(fyne.Size{
		Width:500,
		Height:500,
	})
	w.ShowAndRun()

}


