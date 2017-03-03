package utils

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/multimedia"
)

type Player struct {
	*multimedia.QMediaPlayer
	Model
}

func NewPlayer() *Player {
	player := Player{}
	InitModel(&player)
	player.init()
	return &player
}

func (this *Player) init() {
	this.QMediaPlayer = multimedia.NewQMediaPlayer(nil, 0)
}

func (this *Player) SetFileName(fileName string) {
	url := core.QUrl_FromLocalFile(fileName)
	this.Log.Debug("%v", url.IsValid())
	mediaContent := multimedia.NewQMediaContent2(url)
	this.Log.Debug("%v", mediaContent.IsNull())
	this.QMediaPlayer.SetMedia(mediaContent, nil)
}

func (this *Player) SetVolume(volume int) {
	this.QMediaPlayer.SetVolume(volume)
}

func (this *Player) Play() {
	this.QMediaPlayer.Play()
}

func (this *Player) Pause() {
	this.QMediaPlayer.Pause()
}

func (this *Player) Stop() {
	this.QMediaPlayer.Stop()
}
