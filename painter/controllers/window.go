package controllers

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"painter/models"
	"painter/views"
)

type MainWindow struct {
	*widgets.QMainWindow
	Model
	topToolFrame *views.TopToolFrame
	painterFrame *views.PainterFrame
	painter      *models.Painter
}

func NewMainWindow() *MainWindow {
	window := MainWindow{}
	InitModel(&window)
	window.init()
	return &window
}

func (this *MainWindow) init() {
	this.QMainWindow = widgets.NewQMainWindow(nil, 0)
	this.SetWindowTitle("Painter")
	this.Resize2(800, 600)

	this.topToolFrame = views.NewTopToolFrame(this)
	this.painterFrame = views.NewPainterFrame(this)
	this.painter = models.NewPainter()

	this.topToolFrame.SetOnShapeClickListener(func(id string) {
		this.painter.SetShape(id)
	})
	this.topToolFrame.SetOnColorClickListener(func(id string) {
		this.painter.SetColor(id)
	})
	this.painterFrame.SetDrawListener(func(begin *core.QPoint, end *core.QPoint, painter *gui.QPainter) {
		x1 := begin.X()
		y1 := begin.Y()
		x2 := end.X()
		y2 := end.Y()
		width := x2 - x1
		height := y2 - y1

		color := this.painter.GetColor()
		shape := this.painter.GetShape()

		penColor := map[string]*gui.QColor{
			"red":   gui.NewQColor2(core.Qt__red),
			"green": gui.NewQColor2(core.Qt__green),
			"blue":  gui.NewQColor2(core.Qt__blue),
		}
		painter.SetPen2(penColor[color])

		if shape == "line" {
			painter.DrawLine3(x1, y1, x2, y2)
		} else if shape == "circle" {
			painter.DrawEllipse3(x1, y1, width, height)
		} else {
			painter.DrawRect2(x1, y1, width, height)
		}

	})
}
