package views

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type MusicLrcList struct {
	*widgets.QListView
	Model
	parent widgets.QWidget_ITF
	model  *gui.QStandardItemModel
}

func NewMusicLrcList(parent widgets.QWidget_ITF) *MusicLrcList {
	musicLrcList := MusicLrcList{}
	InitModel(&musicLrcList)
	musicLrcList.init(parent)
	return &musicLrcList
}

func (this *MusicLrcList) init(parent widgets.QWidget_ITF) {
	this.parent = parent
	this.QListView = widgets.NewQListView(parent)
	this.model = gui.NewQStandardItemModel(nil)

	this.SetModel(this.model)
	this.SetWordWrap(true)
	this.SetUniformItemSizes(true)
	this.SetGridSize(core.NewQSize2(-1, 50))
	this.SetFont(gui.NewQFont2("Microsoft YaHei", 15, -1, false))
	this.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	this.SetFocusPolicy(core.Qt__NoFocus)
	this.SetSelectionMode(widgets.QAbstractItemView__NoSelection)
	this.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	this.SetAcceptDrops(true)
}

func (this *MusicLrcList) SetLrc(lrc []string) {
	this.model.Clear()
	for _, singleLrc := range lrc {
		item := gui.NewQStandardItem2(singleLrc)
		item.SetTextAlignment(core.Qt__AlignCenter)
		item.SetFont(gui.NewQFont2("Microsoft YaHei", -1, 50, false))
		this.model.AppendRow2(item)
	}
}
