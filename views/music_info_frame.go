package views

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type MusicInfoFrame struct {
	*widgets.QFrame
	Model
	parent       widgets.QWidget_ITF
	titleLabel   *widgets.QLabel
	artistLabel  *widgets.QLabel
	musicLrcList *MusicLrcList
}

func NewMusicInfoFrame(parent widgets.QWidget_ITF) *MusicInfoFrame {
	musicInfoFrame := MusicInfoFrame{}
	InitModel(&musicInfoFrame)
	musicInfoFrame.init(parent)
	return &musicInfoFrame
}

func (this *MusicInfoFrame) init(parent widgets.QWidget_ITF) {
	this.parent = parent
	this.QFrame = widgets.NewQFrame(parent, 0)
	this.SetGeometry2(0, 60, 800-300, 600-120)

	this.titleLabel = widgets.NewQLabel(nil, 0)
	this.titleLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Fixed)
	this.titleLabel.SetMaximumWidth(800 - 300)
	this.titleLabel.SetWordWrap(true)
	titleFont := gui.NewQFont2("Microsoft YaHei", 20, 75, false)
	this.titleLabel.SetFont(titleFont)
	this.titleLabel.SetAlignment(core.Qt__AlignCenter)

	this.artistLabel = widgets.NewQLabel(nil, 0)
	this.artistLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Fixed)
	this.artistLabel.SetMaximumWidth(800 - 300)
	this.artistLabel.SetAlignment(core.Qt__AlignCenter)

	this.musicLrcList = NewMusicLrcList(parent)

	vLayout := widgets.NewQVBoxLayout2(this)
	vLayout.AddWidget(this.titleLabel, 0, 0)
	vLayout.AddWidget(this.artistLabel, 0, 0)
	vLayout.AddWidget(this.musicLrcList, 0, 0)
	vLayout.AddStretch(0)
	vLayout.SetStretchFactor(this.titleLabel, 1)
	vLayout.SetStretchFactor(this.artistLabel, 1)
	vLayout.SetStretchFactor(this.musicLrcList, 10)
}

func (this *MusicInfoFrame) SetTitle(title string) {
	this.titleLabel.SetText(title)
}

func (this *MusicInfoFrame) SetArtist(artist string) {
	this.artistLabel.SetText(artist)
}

func (this *MusicInfoFrame) SetLrc(lrc []string) {
	this.musicLrcList.SetLrc(lrc)
}
