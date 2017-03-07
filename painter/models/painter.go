package models

import (
	. "github.com/fishedee/web"
)

type Painter struct {
	Model
	shape string
	color string
}

func NewPainter() *Painter {
	painter := Painter{}
	InitModel(&painter)
	painter.init()
	return &painter
}

func (this *Painter) init() {
	this.shape = "line"
	this.color = "red"
}

func (this *Painter) SetShape(shape string) {
	this.shape = shape
}

func (this *Painter) GetShape() string {
	return this.shape
}

func (this *Painter) SetColor(color string) {
	this.color = color
}

func (this *Painter) GetColor() string {
	return this.color
}
