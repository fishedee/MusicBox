package views

import (
	. "github.com/fishedee/web"
	//"github.com/therecipe/qt/core"
	//"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type TopToolFrameClickListener func(id string)

type TopToolFrame struct {
	*widgets.QFrame
	Model
	parent       widgets.QWidget_ITF
	lineButton   *widgets.QRadioButton
	rectButton   *widgets.QRadioButton
	circleButton *widgets.QRadioButton
	redButton    *widgets.QRadioButton
	greenButton  *widgets.QRadioButton
	blueButton   *widgets.QRadioButton
}

func NewTopToolFrame(parent widgets.QWidget_ITF) *TopToolFrame {
	topToolFrame := TopToolFrame{}
	InitModel(&topToolFrame)
	topToolFrame.init(parent)
	return &topToolFrame
}

func (this *TopToolFrame) init(parent widgets.QWidget_ITF) {
	this.QFrame = widgets.NewQFrame(parent, 0)
	this.SetGeometry2(0, 0, 800, 60)

	this.lineButton = widgets.NewQRadioButton(nil)
	this.lineButton.SetText("直线")

	this.rectButton = widgets.NewQRadioButton(nil)
	this.rectButton.SetText("矩形")

	this.circleButton = widgets.NewQRadioButton(nil)
	this.circleButton.SetText("圆形")

	this.redButton = widgets.NewQRadioButton(nil)
	this.redButton.SetText("红色")

	this.greenButton = widgets.NewQRadioButton(nil)
	this.greenButton.SetText("绿色")

	this.blueButton = widgets.NewQRadioButton(nil)
	this.blueButton.SetText("蓝色")

	btnGroup1 := widgets.NewQButtonGroup(this)
	btnGroup1.AddButton(this.lineButton, 0)
	btnGroup1.AddButton(this.rectButton, 0)
	btnGroup1.AddButton(this.circleButton, 0)

	btnGroup2 := widgets.NewQButtonGroup(this)
	btnGroup2.AddButton(this.redButton, 0)
	btnGroup2.AddButton(this.greenButton, 0)
	btnGroup2.AddButton(this.blueButton, 0)

	hLayout := widgets.NewQHBoxLayout2(this)
	hLayout.AddWidget(this.lineButton, 0, 0)
	hLayout.AddWidget(this.rectButton, 0, 0)
	hLayout.AddWidget(this.circleButton, 0, 0)
	hLayout.AddWidget(this.redButton, 0, 0)
	hLayout.AddWidget(this.greenButton, 0, 0)
	hLayout.AddWidget(this.blueButton, 0, 0)

	this.lineButton.SetChecked(true)
	this.redButton.SetChecked(true)
}

func (this *TopToolFrame) SetOnShapeClickListener(listener TopToolFrameClickListener) {
	buttons := map[string]*widgets.QRadioButton{
		"line":      this.lineButton,
		"rectangle": this.rectButton,
		"circle":    this.circleButton,
	}
	for name, button := range buttons {
		buttonName := name
		button.ConnectClicked(func(bool) {
			listener(buttonName)
		})
	}
}

func (this *TopToolFrame) SetOnColorClickListener(listener TopToolFrameClickListener) {
	buttons := map[string]*widgets.QRadioButton{
		"red":   this.redButton,
		"green": this.greenButton,
		"blue":  this.blueButton,
	}
	for name, button := range buttons {
		buttonName := name
		button.ConnectClicked(func(bool) {
			listener(buttonName)
		})
	}
}
