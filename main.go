package main

import (
	. "github.com/fishedee/web"
	"os"
    "github.com/therecipe/qt/widgets"
)

type MainApplication struct{
	Model
}

func (this *MainApplication) Go(){
	widgets.NewQApplication(len(os.Args), os.Args)
	window := NewMainWindow()
	window.Show()
	widgets.QApplication_Exec()
}

func main(){
	mainApplication := MainApplication{}
	InitModel(&mainApplication)
	mainApplication.Go()
}