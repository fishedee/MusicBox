package models

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/multimedia"
	"time"
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
	go func() {
		time.Sleep(3 * time.Second)
		this.Log.Debug("%v", this.QMediaPlayer.State())
		this.Log.Debug("%v", this.QMediaPlayer.Volume())
		this.Log.Debug("%v", this.QMediaPlayer.Error())
		this.Log.Debug("%v", this.QMediaPlayer.MediaStatus())
	}()

}

func (this *Player) Pause() {
	this.QMediaPlayer.Pause()
}

func (this *Player) Stop() {
	this.QMediaPlayer.Stop()
}
