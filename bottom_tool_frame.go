package main

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type BottomToolFrame struct {
	*widgets.QFrame
	Model
	parent       widgets.QWidget_ITF
	lastButton   *widgets.QPushButton
	playButton   *widgets.QPushButton
	stopButton   *widgets.QPushButton
	nextButton   *widgets.QPushButton
	orderButton  *widgets.QPushButton
	seekSlider   *widgets.QSlider
	timeLabel    *widgets.QLabel
	volumnLabel  *widgets.QLabel
	volumnSlider *widgets.QSlider
	curTimeInt   int
}

func NewBottomToolFrame(parent widgets.QWidget_ITF) *BottomToolFrame {
	bottomToolFrame := BottomToolFrame{}
	InitModel(&bottomToolFrame)
	bottomToolFrame.init(parent)
	return &bottomToolFrame
}

func (this *BottomToolFrame) init(parent widgets.QWidget_ITF) {
	this.QFrame = widgets.NewQFrame(parent, 0)
	this.SetStyleSheet(`
		BottomToolFrame{
        border-width: 1px 0 0 0;
        border-style: solid;
        border-color: gray;
        }
	`)

	//rect := parent.Rect()
	this.SetGeometry2(0, 600-60, 800, 60)

	this.lastButton = widgets.NewQPushButton(nil)
	this.lastButton.SetText("上一首")

	this.playButton = widgets.NewQPushButton(nil)
	this.playButton.SetText("播放")

	this.stopButton = widgets.NewQPushButton(nil)
	this.stopButton.SetText("停止")
	this.stopButton.SetEnabled(false)

	this.nextButton = widgets.NewQPushButton(nil)
	this.nextButton.SetText("下一首")

	this.orderButton = widgets.NewQPushButton(nil)
	this.orderButton.SetText("顺序播放")

	this.seekSlider = widgets.NewQSlider2(core.Qt__Horizontal, nil)

	this.timeLabel = widgets.NewQLabel2("00:00/00:00", nil, 0)
	this.timeLabel.SetAlignment(core.Qt__AlignCenter)
	this.timeLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Fixed)

	this.volumnLabel = widgets.NewQLabel(nil, 0)
	this.volumnLabel.SetPixmap(gui.NewQPixmap5("res/volumn.png", "", 0))
	this.volumnSlider = widgets.NewQSlider2(core.Qt__Horizontal, nil)

	hLayout := widgets.NewQHBoxLayout2(this)
	hLayout.AddWidget(this.lastButton, 0, 0)
	hLayout.AddWidget(this.playButton, 0, 0)
	hLayout.AddWidget(this.stopButton, 0, 0)
	hLayout.AddWidget(this.nextButton, 0, 0)
	hLayout.AddWidget(this.orderButton, 0, 0)
	hLayout.AddWidget(this.seekSlider, 0, 0)
	hLayout.AddWidget(this.timeLabel, 0, 0)
	hLayout.AddWidget(this.volumnLabel, 0, 0)
	hLayout.AddWidget(this.volumnSlider, 0, 0)
	hLayout.AddStretch(0)
	hLayout.SetStretchFactor(this.lastButton, 1)
	hLayout.SetStretchFactor(this.playButton, 1)
	hLayout.SetStretchFactor(this.stopButton, 1)
	hLayout.SetStretchFactor(this.nextButton, 1)
	hLayout.SetStretchFactor(this.orderButton, 1)
	hLayout.SetStretchFactor(this.seekSlider, 10)
	hLayout.SetStretchFactor(this.timeLabel, 1)
	hLayout.SetStretchFactor(this.volumnLabel, 1)
	hLayout.SetStretchFactor(this.volumnSlider, 5)

	//FIXME
	this.lastButton.ConnectClicked(this.fixme)
	this.playButton.ConnectClicked(this.fixme)
	this.stopButton.ConnectClicked(this.fixme)
	this.nextButton.ConnectClicked(this.fixme)
	this.orderButton.ConnectClicked(this.fixme)
}

func (this *BottomToolFrame) fixme(checked bool) {
	//FIXME
}
