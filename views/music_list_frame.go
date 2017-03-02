package views

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/widgets"
)

type MusicListFrame struct {
	*widgets.QTabWidget
	Model
	parent        widgets.QWidget_ITF
	allMusicTable *MusicListTable
	favMusicTable *MusicListTable
}

func NewMusicListFrame(parent widgets.QWidget_ITF) *MusicListFrame {
	musicListFrame := MusicListFrame{}
	InitModel(&musicListFrame)
	musicListFrame.init(parent)
	return &musicListFrame
}

func (this *MusicListFrame) init(parent widgets.QWidget_ITF) {
	this.parent = parent
	this.QTabWidget = widgets.NewQTabWidget(parent)
	this.allMusicTable = NewMusicListTable(nil)
	this.favMusicTable = NewMusicListTable(nil)

	this.SetGeometry2(800-300, 60, 301, 600-117)
	this.AddTab(this.allMusicTable, "播放列表")
	this.AddTab(this.favMusicTable, "我的收藏")
	this.SetCurrentIndex(0)

	this.allMusicTable.SetContextMenuListener(func(index int) []MusicListContextAction {
		this.Log.Debug("%v", index)
		return []MusicListContextAction{
			MusicListContextAction{
				Name: "播放",
				Action: func(actionRows []int) {

				},
			},
			MusicListContextAction{
				Name:   "",
				Action: nil,
			},
			MusicListContextAction{
				Name: "移除",
				Action: func(actionRows []int) {

				},
			},
			MusicListContextAction{
				Name:   "",
				Action: nil,
			},
			MusicListContextAction{
				Name: "收藏",
				Action: func(actionRows []int) {

				},
			},
			MusicListContextAction{
				Name:   "",
				Action: nil,
			},
			MusicListContextAction{
				Name: "打开文件",
				Action: func(actionRows []int) {

				},
			},
		}
	})
}

func (this *MusicListFrame) AddAllSong(title string, artist string, timeString string) {
	this.allMusicTable.AddSong(title, artist, timeString)
}

func (this *MusicListFrame) AddFavSong(title string, artist string, timeString string) {
	this.favMusicTable.AddSong(title, artist, timeString)
}
