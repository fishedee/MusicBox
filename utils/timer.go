package utils

import (
	"github.com/fishedee/web"
	"github.com/therecipe/qt/core"
)

type Timer struct {
	web.Model
}

func NewTimer() *Timer {
	timer := Timer{}
	web.InitModel(&timer)
	return &timer
}

func (this *Timer) Sleep(parent core.QObject_ITF, milliSecond int, handler func()) {
	timer := core.NewQTimer(parent)
	timer.SetSingleShot(true)
	timer.ConnectTimeout(handler)
	timer.Start(milliSecond)
}
