package main

import (
	"testing"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
)

func Test_make_ui(t *testing.T) {
	var test_cfg config

	edit,preview := test_cfg.make_ui()

	test.Type(edit,"Hello")

	if preview.String() != "Hello" {
		t.Errorf("Failed -- did not find expected value in preview")
	}
}

func Test_RunAPP(t *testing.T){
	var test_cfg  config
	test_app := test.NewApp()
	test_win := test_app.NewWindow("Test MarkDown")

	edit,preview := test_cfg.make_ui()

	test_cfg.create_menu_items(test_win)

	test_win.SetContent(container.NewHSplit(edit,preview))

	test_app.Run()

	test.Type(edit,"Some text")
	if preview.String() != "Some text"{
		t.Error("failed")
	}
	
}
