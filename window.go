package main

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type MainWindow struct {
	*widgets.QMainWindow
	Model
	topToolFrame    *TopToolFrame
	bottomToolFrame *BottomToolFrame
}

func NewMainWindow() *MainWindow {
	window := MainWindow{}
	InitModel(&window)
	window.init()
	return &window
}

func (this *MainWindow) init() {
	this.QMainWindow = widgets.NewQMainWindow(nil, 0)
	icon := gui.NewQIcon5("res/icon.png")
	this.SetWindowIcon(icon)
	this.SetWindowTitle("SxPlayer")
	this.Resize2(800, 600)
	this.topToolFrame = NewTopToolFrame(this)
	this.bottomToolFrame = NewBottomToolFrame(this)
}
