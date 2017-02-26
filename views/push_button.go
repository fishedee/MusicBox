package views

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type PushButton struct {
	*widgets.QPushButton
	Model
	parent    widgets.QWidget_ITF
	status    int
	pixmap    *gui.QPixmap
	btnWidth  int
	btnHeight int
}

func NewPushButton(parent widgets.QWidget_ITF) *PushButton {
	pushButton := PushButton{}
	InitModel(&pushButton)
	pushButton.init(parent)
	return &pushButton
}

func (this *PushButton) init(parent widgets.QWidget_ITF) {
	this.QPushButton = widgets.NewQPushButton(parent)
	this.parent = parent
	this.status = 1
	this.ConnectMousePressEvent(this.mousePressEvent)
	this.ConnectMouseReleaseEvent(this.mouseReleaseEvent)
	this.ConnectEnterEvent(this.enterEvent)
	this.ConnectLeaveEvent(this.leaveEvent)
	this.ConnectPaintEvent(this.paintEvent)
}

func (this *PushButton) LoadPixmap(picName string) {
	this.pixmap = gui.NewQPixmap5(picName, "", 0)
	this.btnWidth = this.pixmap.Width() / 4
	this.btnHeight = this.pixmap.Height()
}

func (this *PushButton) enterEvent(event *core.QEvent) {
	if !this.IsChecked() && this.IsEnabled() {
		this.status = 0
		this.Update()
	}
}

func (this *PushButton) leaveEvent(event *core.QEvent) {
	if !this.IsChecked() && this.IsEnabled() {
		this.status = 1
		this.Update()
	}
}

func (this *PushButton) SetDisabled(disabled bool) {
	this.QPushButton.SetDisabled(disabled)
	if !this.IsEnabled() {
		this.status = 2
		this.Update()
	} else {
		this.status = 1
		this.Update()
	}
}

func (this *PushButton) mousePressEvent(event *gui.QMouseEvent) {
	if event.Button() == core.Qt__LeftButton {
		this.status = 2
		this.Update()
	}
}

func (this *PushButton) mouseReleaseEvent(event *gui.QMouseEvent) {
	if event.Button() == core.Qt__LeftButton {
		this.Clicked(true)
	}
	if !this.IsChecked() {
		this.status = 3
	}
	if this.Menu() != nil {
		this.Menu().Exec2(event.GlobalPos(), nil)
	}
	this.Update()
}

func (this *PushButton) paintEvent(event *gui.QPaintEvent) {
	painter := gui.NewQPainter()
	painter.Begin(this)
	painter.DrawPixmap10(this.Rect(), this.pixmap.Copy2(this.btnWidth*this.status, 0, this.btnWidth, this.btnHeight))
	painter.End()
}