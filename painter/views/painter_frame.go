package views

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type PainterFrameDrawListener func(begin *core.QPoint, end *core.QPoint, painter *gui.QPainter)
type PainterFrame struct {
	*widgets.QLabel
	Model
	pixmap       *gui.QPixmap
	nextPixmap   *gui.QPixmap
	drawListener PainterFrameDrawListener
	dragPosition *core.QPoint
	isDraging    bool
}

func NewPainterFrame(parent widgets.QWidget_ITF) *PainterFrame {
	painterFrame := PainterFrame{}
	InitModel(&painterFrame)
	painterFrame.init(parent)
	return &painterFrame
}

func (this *PainterFrame) init(parent widgets.QWidget_ITF) {
	this.QLabel = widgets.NewQLabel(parent, 0)
	this.SetGeometry2(0, 60, 800, 600-60)

	this.pixmap = gui.NewQPixmap3(800, 600-60)
	this.pixmap.Fill(gui.NewQColor2(core.Qt__white))
	this.SetPixmap(this.pixmap)
	this.ConnectMousePressEvent(this.pressEvent)
	this.ConnectMouseReleaseEvent(this.releaseEvent)
	this.ConnectMouseMoveEvent(this.moveEvent)
}

func (this *PainterFrame) SetDrawListener(listener PainterFrameDrawListener) {
	this.drawListener = listener
}

func (this *PainterFrame) pressEvent(event *gui.QMouseEvent) {
	this.isDraging = true
	this.dragPosition = event.Pos()
	this.nextPixmap = nil
}

func (this *PainterFrame) releaseEvent(event *gui.QMouseEvent) {
	this.isDraging = false
	if this.nextPixmap != nil {
		this.pixmap = this.nextPixmap
	}
}

func (this *PainterFrame) moveEvent(event *gui.QMouseEvent) {
	if this.isDraging {
		nowDragPosition := event.Pos()
		if this.drawListener != nil {
			this.nextPixmap = gui.NewQPixmap3(800, 600-60)
			painter := gui.NewQPainter()
			painter.Begin(this.nextPixmap)
			painter.DrawPixmap9(0, 0, this.pixmap)
			this.drawListener(this.dragPosition, nowDragPosition, painter)
			painter.End()
			this.SetPixmap(this.nextPixmap)
		}
	}
}
