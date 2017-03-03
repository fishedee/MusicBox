package models

import (
	. "github.com/fishedee/web"
)

type Music struct {
	Id         int
	FilePath   string
	FileName   string
	FileFormat string
	Title      string
	Artist     string
}

type MusicList struct {
	Model
	allMusic  []Music
	favMusic  []Music
	playIsAll bool
	playIndex int
	idCounter int
	listener  func(music Music)
}

func NewMusicList() *MusicList {
	musicList := MusicList{}
	InitModel(&musicList)
	musicList.init()
	return &musicList
}

func (this *MusicList) init() {
	this.allMusic = []Music{}
	this.favMusic = []Music{}
	this.playIsAll = true
	this.playIndex = -1
	this.idCounter = 10001
}

func (this *MusicList) AddAllMusic(music Music) {
	music.Id = this.idCounter
	this.idCounter++
	this.allMusic = append(this.allMusic, music)
	if len(this.allMusic) == 1 {
		this.play(true, 0)
	}
}

func (this *MusicList) DelAllMusic(index int) {
	if index < 0 || index >= len(this.allMusic) {
		return
	}
	delMusic := this.allMusic[index]
	this.UnFavAllMusic(delMusic.Id)
	newMusic := []Music{}
	for singleIndex, singleMusic := range this.allMusic {
		if singleIndex == index {
			continue
		}
		newMusic = append(newMusic, singleMusic)
	}
	this.allMusic = newMusic
	if this.playIsAll {
		if this.playIndex > index {
			this.playIndex--
		} else if this.playIndex == index {
			if len(this.allMusic) == 0 {
				this.playIndex = -1
			} else {
				this.play(true, this.playIndex)
			}
		}
	}
}

func (this *MusicList) PlayAllMusic(index int) {
	if index < 0 || index >= len(this.allMusic) {
		return
	}
	this.play(true, index)
}

func (this *MusicList) getFavAllMusic(id int) int {
	for singleIndex, singleMusic := range this.favMusic {
		if singleMusic.Id == id {
			return singleIndex
		}
	}
	return -1
}

func (this *MusicList) IsFavAllMusic(index int) bool {
	if index < 0 || index >= len(this.allMusic) {
		return false
	}
	return this.getFavAllMusic(this.allMusic[index].Id) != -1
}

func (this *MusicList) FavAllMusic(index int) {
	if index < 0 || index >= len(this.allMusic) {
		return
	}
	favSingleMusic := this.allMusic[index]
	this.favMusic = append(this.favMusic, favSingleMusic)
}

func (this *MusicList) UnFavAllMusic(index int) {
	if index < 0 || index >= len(this.allMusic) {
		return
	}
	favMusicIndex := this.getFavAllMusic(this.allMusic[index].Id)
	if favMusicIndex == -1 {
		return
	}
	this.DelFavMusic(favMusicIndex)
}

func (this *MusicList) DelFavMusic(index int) {
	if index < 0 || index >= len(this.favMusic) {
		return
	}
	newMusic := []Music{}
	for singleIndex, singleMusic := range this.favMusic {
		if singleIndex == index {
			continue
		}
		newMusic = append(newMusic, singleMusic)
	}
	this.favMusic = newMusic
	if this.playIsAll == false {
		if this.playIndex > index {
			this.playIndex--
		} else if this.playIndex == index {
			if len(this.favMusic) == 0 {
				this.playIndex = -1
			} else {
				this.play(false, this.playIndex)
			}
		}
	}
}

func (this *MusicList) PlayFavMusic(index int) {
	if index < 0 || index >= len(this.favMusic) {
		return
	}
	this.play(false, index)
}

func (this *MusicList) Prev() {
	var music []Music
	if this.playIsAll {
		music = this.allMusic
	} else {
		music = this.favMusic
	}
	musicLen := len(music)
	if musicLen == 0 {
		return
	}
	this.playIndex = (this.playIndex - 1 + musicLen) % musicLen
}

func (this *MusicList) Next() {
	var music []Music
	if this.playIsAll {
		music = this.allMusic
	} else {
		music = this.favMusic
	}
	musicLen := len(music)
	if musicLen == 0 {
		return
	}
	this.playIndex = (this.playIndex + 1) % musicLen
}

func (this *MusicList) play(playIsAll bool, playIndex int) {
	this.playIsAll = playIsAll
	this.playIndex = playIndex
	var music []Music
	if this.playIsAll {
		music = this.allMusic
	} else {
		music = this.favMusic
	}
	if this.listener != nil {
		this.listener(music[this.playIndex])
	}
}

func (this *MusicList) SetPlayListener(listener func(music Music)) {
	this.listener = listener
}
