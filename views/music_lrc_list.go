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
	parent   widgets.QWidget_ITF
	model    *gui.QStandardItemModel
	curIndex int
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

	this.curIndex = -1
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

func (this *MusicLrcList) ActiveIndex(index int) {
	if this.curIndex != -1 {
		color := gui.NewQColor3(0, 0, 0, 0)
		brush := gui.NewQBrush3(color, 0)
		this.model.Item(this.curIndex, 0).SetForeground(brush)
		this.model.Item(this.curIndex, 1).SetForeground(brush)
		this.model.Item(this.curIndex, 2).SetForeground(brush)
	}
	if index != -1 {
		color := gui.NewQColor3(255, 0, 0, 0)
		brush := gui.NewQBrush3(color, 0)
		this.model.Item(index, 0).SetForeground(brush)
		this.model.Item(index, 1).SetForeground(brush)
		this.model.Item(index, 2).SetForeground(brush)
		this.ScrollTo(this.model.Index(index, 0, nil), widgets.QAbstractItemView__PositionAtCenter)
	}
	this.curIndex = index
}
