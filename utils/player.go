package utils

import (
	"fmt"
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
	mediaContent := multimedia.NewQMediaContent2(url)
	this.QMediaPlayer.SetMedia(mediaContent, nil)
}

func (this *Player) GetPosition() (int, int, int, string) {
	minPosition := 0
	maxPosition := int(this.QMediaPlayer.Duration() / 1000)
	curPosition := int(this.QMediaPlayer.Position() / 1000)

	positionString := fmt.Sprintf("%02d:%02d/%02d:%02d", curPosition/60, curPosition%60, maxPosition/60, maxPosition%60)
	return minPosition, maxPosition, curPosition, positionString
}

func (this *Player) SetPosition(position int) {
	this.QMediaPlayer.SetPosition(int64(position) * 1000)
}

func (this *Player) GetVolume() (int, int, int) {
	return 0, 100, this.QMediaPlayer.Volume()
}

func (this *Player) SetVolume(volume int) {
	this.QMediaPlayer.SetVolume(volume)
}

func (this *Player) GetError() (int, string) {
	return int(this.QMediaPlayer.Error()), this.QMediaPlayer.ErrorString()
}

func (this *Player) GetMetaData() map[string]string {
	title := this.QMediaPlayer.MetaData("Title").ToString()
	subTitle := this.QMediaPlayer.MetaData("SubTitle").ToString()
	author := this.QMediaPlayer.MetaData("Author").ToString()
	lyrics := this.QMediaPlayer.MetaData("Lyrics").ToString()
	duration := this.QMediaPlayer.MetaData("Duration").ToString()
	albumTitle := this.QMediaPlayer.MetaData("AlbumTitle").ToString()
	albumArtist := this.QMediaPlayer.MetaData("AlbumArtist").ToString()
	contributingArtist := this.QMediaPlayer.MetaData("ContributingArtist").ToString()
	if author == "" {
		if contributingArtist != "" {
			author = contributingArtist
		}
	}
	return map[string]string{
		"title":       title,
		"subTitle":    subTitle,
		"author":      author,
		"lyrics":      lyrics,
		"duration":    duration,
		"albumTitle":  albumTitle,
		"albumArtist": albumArtist,
	}
}

func (this *Player) IsPause() bool {
	return this.State() == multimedia.QMediaPlayer__PausedState
}

func (this *Player) IsPlay() bool {
	return this.State() == multimedia.QMediaPlayer__PlayingState
}

func (this *Player) IsStop() bool {
	return this.State() == multimedia.QMediaPlayer__StoppedState
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

func (this *Player) SetMetaListener(handler func()) {
	this.QMediaPlayer.ConnectMetaDataAvailableChanged(func(available bool) {
		handler()
	})
}

func (this *Player) SetErrorListener(handler func()) {
	this.QMediaPlayer.ConnectMediaStatusChanged(func(status multimedia.QMediaPlayer__MediaStatus) {
		if status == multimedia.QMediaPlayer__InvalidMedia {
			handler()
		}
	})
}

func (this *Player) SetPositionChangeListener(handler func()) {
	this.QMediaPlayer.ConnectPositionChanged(func(position int64) {
		handler()
	})
}

func (this *Player) SetDurationChangeListener(handler func()) {
	this.QMediaPlayer.ConnectDurationChanged(func(duration int64) {
		handler()
	})
}

func (this *Player) SetEndListener(handler func()) {
	this.QMediaPlayer.ConnectMediaStatusChanged(func(status multimedia.QMediaPlayer__MediaStatus) {
		if status == multimedia.QMediaPlayer__EndOfMedia {
			handler()
		}
	})
}
