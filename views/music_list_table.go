package views

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type MusicListTable struct {
	*widgets.QTableView
	Model
	parent widgets.QWidget_ITF
	model  *gui.QStandardItemModel
}

type MusicListContextAction struct {
	Name   string
	Action func(actionRows []int)
}

func NewMusicListTable(parent widgets.QWidget_ITF) *MusicListTable {
	musicListTable := MusicListTable{}
	InitModel(&musicListTable)
	musicListTable.init(parent)
	return &musicListTable
}

func (this *MusicListTable) init(parent widgets.QWidget_ITF) {
	this.parent = parent
	this.QTableView = widgets.NewQTableView(parent)
	this.model = gui.NewQStandardItemModel(parent)
	this.model.SetColumnCount(3)
	this.SetFixedWidth(300)
	this.SetShowGrid(false)
	this.SetWordWrap(false)
	this.SetMouseTracking(true)

	this.VerticalHeader().SetSectionResizeMode(widgets.QHeaderView__ResizeToContents)
	this.VerticalHeader().SetSectionsClickable(false)

	this.HorizontalHeader().SetSectionResizeMode(widgets.QHeaderView__Stretch)
	this.HorizontalHeader().SetSectionsClickable(false)
	this.HorizontalHeader().Hide()
	this.HorizontalHeader().SetStyleSheet(`
        selection-background-color:lightblue;
    `)

	this.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)
	this.SetSelectionMode(widgets.QAbstractItemView__ExtendedSelection)

	this.SetFocusPolicy(core.Qt__NoFocus)
	this.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	this.SetStyleSheet(`
        selection-background-color:lightblue;
    `)

	this.SetModel(this.model)

}

func (this *MusicListTable) AddSong(title string, artist string, timeString string) {
	song := []*gui.QStandardItem{
		gui.NewQStandardItem2(title),
		gui.NewQStandardItem2(artist),
		gui.NewQStandardItem2(timeString),
	}
	song[2].SetTextAlignment(core.Qt__AlignCenter)
	this.model.AppendRow(song)
}

func (this *MusicListTable) DelSong(index int) {
	this.model.RemoveRow(index, nil)
}

func (this *MusicListTable) ClearSong() {
	this.model.Clear()
}

func (this *MusicListTable) SetDoubleClickListener(handler func(index int)) {
	this.ConnectMouseDoubleClickEvent(func(event *gui.QMouseEvent) {
		if event.Button() == core.Qt__LeftButton {
			pos := event.Pos()
			index := this.IndexAt(pos).Row()
			handler(index)
		}
	})
}

func (this *MusicListTable) ActiveIndex(index int) {
	color := gui.NewQColor3(0, 0, 0, 0)
	brush := gui.NewQBrush3(color, 0)
	this.model.Item(index, 0).SetForeground(brush)
	this.model.Item(index, 1).SetForeground(brush)
	this.model.Item(index, 2).SetForeground(brush)
}

func (this *MusicListTable) getContextMenu(actions []MusicListContextAction) *widgets.QMenu {
	contextMenu := widgets.NewQMenu(this.parent)
	for _, singleAction := range actions {
		if singleAction.Name != "" {
			action := widgets.NewQAction2(singleAction.Name, this.parent)
			action.ConnectTrigger(func() {
				selectedIndexs := this.SelectionModel().SelectedRows(0)
				selectedRows := []int{}
				for _, singleIndex := range selectedIndexs {
					row := this.model.ItemFromIndex(singleIndex).Row()
					selectedRows = append(selectedRows, row)
				}
				singleAction.Action(selectedRows)
			})
			contextMenu.QWidget.AddAction(action)
		} else {
			contextMenu.AddSeparator()
		}

	}
	return contextMenu
}

func (this *MusicListTable) SetContextMenuListener(handler func(index int) []MusicListContextAction) {
	this.ConnectContextMenuEvent(func(event *gui.QContextMenuEvent) {
		pos := event.Pos()
		index := this.IndexAt(pos).Row()
		actions := handler(index)
		contextMenu := this.getContextMenu(actions)
		contextMenu.Exec2(event.GlobalPos(), nil)
	})
}
