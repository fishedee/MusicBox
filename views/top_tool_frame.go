package views

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type TopToolFrame struct {
	*widgets.QFrame
	Model
	parent    widgets.QWidget_ITF
	isDraging bool
}

func NewTopToolFrame(parent widgets.QWidget_ITF) *TopToolFrame {
	topToolFrame := TopToolFrame{}
	InitModel(&topToolFrame)
	topToolFrame.init(parent)
	return &topToolFrame
}

func (this *TopToolFrame) init(parent widgets.QWidget_ITF) {
	this.QFrame = widgets.NewQFrame(parent, 0)
	this.parent = parent
	this.SetStyleSheet(`border-image:url(res/box.png);
		background-repeat:no-repeat;
	`)
	this.isDraging = false
	this.SetGeometry2(0, 0, 800, 60)
	this.AddRobotLogo()
	this.AddButtons()
	this.ConnectMousePressEvent(this.mousePressEvent)
	this.ConnectMouseReleaseEvent(this.mouseReleaseEvent)
	this.ConnectMouseMoveEvent(this.mouseMoveEvent)
}

func (this *TopToolFrame) AddRobotLogo() {
	btn := widgets.NewQPushButton(this.parent)
	btn.SetObjectName("btnSpecial")
	btn.SetStyleSheet(`
		QPushButton#btnSpecial {
        border-image: url(res/robot_1.png);
        background-repeat: no-repeat;
        }
        QPushButton#btnSpecial:hover {
        border-image: url(res/robot_2.png);
        background-repeat: no-repeat;
        }
        QPushButton#btnSpecial:pressed {
        border-image: url(res/robot_3.png);
        background-repeat: no-repeat;
        }
	`)
	btn.SetGeometry2(20, 0, 67, 60)
}

func (this *TopToolFrame) AddButtons() {
	closeButton := NewPushButton(this)
	closeButton.LoadPixmap("res/close.png")
	closeButton.SetGeometry2(770, 10, 16, 16)
	//FIXME

	miniButton := NewPushButton(this)
	miniButton.LoadPixmap("res/mini.png")
	miniButton.SetGeometry2(740, 10, 16, 16)
	//FIXME
}

func (this *TopToolFrame) mousePressEvent(event *gui.QMouseEvent) {
	this.isDraging = true
	//FIXME
}

func (this *TopToolFrame) mouseReleaseEvent(event *gui.QMouseEvent) {
	this.isDraging = false
	//FIXME
}

func (this *TopToolFrame) mouseMoveEvent(event *gui.QMouseEvent) {
	if this.isDraging {
		//FIXME
	}
}
