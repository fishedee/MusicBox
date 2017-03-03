package controllers

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"musicbox/models"
	"musicbox/utils"
	"musicbox/views"
)

type MainWindow struct {
	*widgets.QMainWindow
	Model
	topToolFrame    *views.TopToolFrame
	bottomToolFrame *views.BottomToolFrame
	musicListFrame  *views.MusicListFrame
	musicInfoFrame  *views.MusicInfoFrame
	player          *utils.Player
}

func NewMainWindow() *MainWindow {
	window := MainWindow{}
	InitModel(&window)
	window.init()
	return &window
}

func (this *MainWindow) init() {
	this.QMainWindow = widgets.NewQMainWindow(nil, 0)
	icon := gui.NewQIcon5("res/icon.png")
	this.SetWindowIcon(icon)
	this.SetWindowTitle("SxPlayer")
	this.Resize2(800, 600)
	this.topToolFrame = views.NewTopToolFrame(this)
	this.bottomToolFrame = views.NewBottomToolFrame(this)
	this.musicListFrame = views.NewMusicListFrame(this)
	this.musicInfoFrame = views.NewMusicInfoFrame(this)

	this.musicListFrame.AddAllSong("标题1", "作者1", "11:11")
	this.musicListFrame.AddAllSong("标题2", "作者2", "11:11")
	this.musicListFrame.AddFavSong("标题3", "作者3", "11:11")

	this.musicInfoFrame.SetTitle("标题1")
	this.musicInfoFrame.SetArtist("作者1")
	this.musicInfoFrame.SetLrc([]string{
		"歌词1",
		"歌词2",
		"歌词3",
	})

	this.player = utils.NewPlayer()
	this.player.SetFileName("/Users/fishedee/Project/fishgo/src/musicbox/res/test.mp3")
	this.player.SetVolume(100)
	this.player.Play()
}
