package views

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type PushButton struct {
	*widgets.QLabel
	Model
	parent    widgets.QWidget_ITF
	status    int
	pixmap    *gui.QPixmap
	btnWidth  int
	btnHeight int
	listener  func()
}

func NewPushButton(parent widgets.QWidget_ITF) *PushButton {
	pushButton := PushButton{}
	InitModel(&pushButton)
	pushButton.init(parent)
	return &pushButton
}

func (this *PushButton) init(parent widgets.QWidget_ITF) {
	this.QLabel = widgets.NewQLabel(parent, 0)
	this.parent = parent
	this.status = 1
	this.ConnectMousePressEvent(this.mousePressEvent)
	this.ConnectMouseReleaseEvent(this.mouseReleaseEvent)
	this.ConnectEnterEvent(this.enterEvent)
	this.ConnectLeaveEvent(this.leaveEvent)
}

func (this *PushButton) LoadPixmap(picName string) {
	this.pixmap = gui.NewQPixmap5(picName, "", 0)
	this.btnWidth = this.pixmap.Size().Width() / 4
	this.btnHeight = this.pixmap.Size().Height()
	this.update()
}

func (this *PushButton) enterEvent(event *core.QEvent) {
	if this.IsEnabled() {
		this.status = 0
		this.update()
	}
}

func (this *PushButton) leaveEvent(event *core.QEvent) {
	if this.IsEnabled() {
		this.status = 1
		this.update()
	}
}

func (this *PushButton) SetDisabled(disabled bool) {
	this.QLabel.SetDisabled(disabled)
	if !this.IsEnabled() {
		this.status = 2
		this.update()
	} else {
		this.status = 1
		this.update()
	}
}

func (this *PushButton) mousePressEvent(event *gui.QMouseEvent) {
	if event.Button() == core.Qt__LeftButton {
		this.status = 2
		this.update()
	}
}

func (this *PushButton) mouseReleaseEvent(event *gui.QMouseEvent) {
	if event.Button() == core.Qt__LeftButton {
		this.status = 0
		if this.listener != nil {
			this.listener()
		}
		this.update()
	}
	this.update()
}

func (this *PushButton) SetClickListener(listener func()) {
	this.listener = listener
}

func (this *PushButton) update() {
	this.SetPixmap(this.pixmap.Copy2(this.btnWidth*this.status, 0, this.btnWidth, this.btnHeight))
}
