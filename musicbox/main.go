package main

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/widgets"
	"musicbox/controllers"
	"os"
)

type MainApplication struct {
	Model
}

func (this *MainApplication) Go() {
	widgets.NewQApplication(len(os.Args), os.Args)
	window := controllers.NewMainWindow()
	window.Show()
	widgets.QApplication_Exec()
}

func main() {
	mainApplication := MainApplication{}
	InitModel(&mainApplication)
	mainApplication.Go()
}
