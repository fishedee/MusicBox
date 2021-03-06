package controllers

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/core"
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
	fileDialog      *utils.FileDialog
	musicList       *models.MusicList
	musicLrc        *models.MusicLrc
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
	this.SetWindowFlags(core.Qt__Window | core.Qt__FramelessWindowHint | core.Qt__WindowMinMaxButtonsHint)
	this.Resize2(800, 600)
	this.topToolFrame = views.NewTopToolFrame(this)
	this.bottomToolFrame = views.NewBottomToolFrame(this)
	this.musicListFrame = views.NewMusicListFrame(this)
	this.musicInfoFrame = views.NewMusicInfoFrame(this)
	this.player = utils.NewPlayer()
	this.fileDialog = utils.NewFileDialog()
	this.musicList = models.NewMusicList()
	this.musicLrc = models.NewMusicLrc()

	this.musicListFrame.SetAllSongContext(func(index int) []views.MusicListContextAction {
		isFav := this.musicList.IsFavAllMusic(index)
		var favAction views.MusicListContextAction
		if isFav {
			favAction = views.MusicListContextAction{
				Name: "取消收藏",
				Action: func(actionRows []int) {
					if len(actionRows) == 0 {
						return
					}
					this.musicList.UnFavAllMusic(actionRows[0])
				},
			}
		} else {
			favAction = views.MusicListContextAction{
				Name: "收藏",
				Action: func(actionRows []int) {
					if len(actionRows) == 0 {
						return
					}
					this.musicList.FavAllMusic(actionRows[0])
				},
			}
		}
		return []views.MusicListContextAction{
			views.MusicListContextAction{
				Name: "播放",
				Action: func(actionRows []int) {
					if len(actionRows) == 0 {
						return
					}
					this.musicList.PlayAllMusic(actionRows[0])
				},
			},
			views.MusicListContextAction{
				Name: "",
			},
			views.MusicListContextAction{
				Name: "移除",
				Action: func(actionRows []int) {
					if len(actionRows) == 0 {
						return
					}
					this.musicList.DelAllMusic(actionRows[0])
				},
			},
			views.MusicListContextAction{
				Name: "",
			},
			favAction,
			views.MusicListContextAction{
				Name: "",
			},
			views.MusicListContextAction{
				Name: "打开文件",
				Action: func(actionRows []int) {
					this.fileDialog.Open(this, func(fileInfo []utils.FileInfo) {
						for _, singleFileInfo := range fileInfo {
							this.musicList.AddAllMusic(models.Music{
								Id:         0,
								FilePath:   singleFileInfo.FilePath,
								FileName:   singleFileInfo.FileName,
								FileFormat: singleFileInfo.FileFormat,
								Title:      singleFileInfo.Title,
								Artist:     singleFileInfo.Artist,
								Duration:   singleFileInfo.Duration,
							})
						}
					})

				},
			},
		}
	})
	this.musicListFrame.SetFavSongContext(func(index int) []views.MusicListContextAction {
		return []views.MusicListContextAction{
			views.MusicListContextAction{
				Name: "播放",
				Action: func(actionRows []int) {
					if len(actionRows) == 0 {
						return
					}
					this.musicList.PlayFavMusic(actionRows[0])
				},
			},
			views.MusicListContextAction{
				Name: "",
			},
			views.MusicListContextAction{
				Name: "移除",
				Action: func(actionRows []int) {
					if len(actionRows) == 0 {
						return
					}
					this.musicList.DelFavMusic(actionRows[0])
				},
			},
		}
	})
	this.musicListFrame.SetAllDoubleClickListener(func(index int) {
		this.musicList.PlayAllMusic(index)
	})
	this.musicListFrame.SetFavDoubleClickListener(func(index int) {
		this.musicList.PlayFavMusic(index)
	})

	this.musicLrc.SetDownloadedListener(func() {
		this.loadLrc()
	})

	this.musicList.SetPlayListener(func(music models.Music, playIsAll bool, index int) {
		this.Log.Debug("mm %v", music)
		this.initBottomToolEmpty()
		this.musicInfoFrame.SetTitle(music.Title)
		this.musicInfoFrame.SetArtist(music.Artist)
		this.player.SetFileName(music.FilePath)
		this.player.Play()
		if playIsAll {
			this.musicListFrame.ActiveAllIndex(index)
			this.musicListFrame.ActiveFavIndex(-1)
		} else {
			this.musicListFrame.ActiveAllIndex(-1)
			this.musicListFrame.ActiveFavIndex(index)
		}
		hasLoad := this.loadLrc()
		if hasLoad == false {
			this.musicLrc.DownloadLrc(this, music)
			this.musicInfoFrame.SetLrc([]string{"下载歌词中..."})
			this.musicInfoFrame.ActiveLrcIndex(0)
		} else {
			this.musicInfoFrame.ActiveLrcIndex(0)
		}
	})
	this.musicList.SetAddAllMusicListener(func(music models.Music) {
		this.musicListFrame.AddAllSong(music.Title, music.Artist, music.Duration)
	})
	this.musicList.SetAddFavMusicListener(func(music models.Music) {
		this.musicListFrame.AddFavSong(music.Title, music.Artist, music.Duration)
	})
	this.musicList.SetDelAllMusicListener(func(index int) {
		this.musicListFrame.DelAllSong(index)
	})
	this.musicList.SetDelFavMusicListener(func(index int) {
		this.musicListFrame.DelFavSong(index)
	})

	this.player = utils.NewPlayer()
	this.player.SetPositionChangeListener(func() {
		minPosition, maxPosition, curPosition, curPositionDesc := this.player.GetPosition()
		this.bottomToolFrame.SetSeek(minPosition, maxPosition, curPosition, curPositionDesc)
		lrcProgress := this.musicLrc.GetLrcProgress(this.musicList.GetPlayMusic(), curPosition)
		if lrcProgress != -1 {
			this.musicInfoFrame.ActiveLrcIndex(lrcProgress)
		}
	})
	this.player.SetDurationChangeListener(func() {
		minPosition, maxPosition, curPosition, curPositionDesc := this.player.GetPosition()
		this.bottomToolFrame.SetSeek(minPosition, maxPosition, curPosition, curPositionDesc)
	})
	this.player.SetErrorListener(func() {
		code, msg := this.player.GetError()
		this.Log.Debug("player error!%v,%v", code, msg)
		this.initBottomToolErr()
	})
	this.player.SetEndListener(func() {
		this.musicList.Next()
	})

	this.bottomToolFrame.SetButtonClickListener("prev", func() {
		this.musicList.Prev()
	})
	this.bottomToolFrame.SetButtonClickListener("play", func() {
		if this.player.IsStop() || this.player.IsPause() {
			this.player.Play()
			this.bottomToolFrame.SetButtonText("play", "暂停")
		} else {
			this.player.Pause()
			this.bottomToolFrame.SetButtonText("play", "播放")
		}
	})
	this.bottomToolFrame.SetButtonClickListener("stop", func() {
		this.player.Stop()
		this.bottomToolFrame.SetButtonText("play", "播放")
	})
	this.bottomToolFrame.SetButtonClickListener("next", func() {
		this.musicList.Next()
	})
	this.bottomToolFrame.SetSeekChangeListener(func(progress int) {
		this.player.SetPosition(progress)
	})
	this.bottomToolFrame.SetVolumeChangeListener(func(volume int) {
		this.player.SetVolume(volume)
	})

	this.topToolFrame.SetCloseListener(func() {
		this.Close()
	})
	this.topToolFrame.SetMiniListener(func() {
		this.ShowMinimized()
	})
	this.topToolFrame.SetMoveListener(func(xMove int, yMove int) {
		pos := this.Pos()
		this.Move2(pos.X()+xMove, pos.Y()+yMove)
	})
	this.initBottomToolErr()
}

func (this *MainWindow) loadLrc() bool {
	curMusic := this.musicList.GetPlayMusic()
	if curMusic.Id == 0 {
		return false
	}
	hasLrc := this.musicLrc.HasLrc(curMusic)
	if hasLrc == false {
		return false
	}
	lrcStatus := this.musicLrc.GetLrcStatus(curMusic)
	if lrcStatus != "" {
		this.musicInfoFrame.SetLrc([]string{lrcStatus})
		this.musicInfoFrame.ActiveLrcIndex(0)
		return true
	}
	lrc := this.musicLrc.GetLrcData(curMusic)
	this.musicInfoFrame.SetLrc(lrc)
	return true
}

func (this *MainWindow) initBottomToolErr() {
	this.musicInfoFrame.SetTitle("")
	this.musicInfoFrame.SetArtist("")
	this.bottomToolFrame.SetButtonText("play", "播放")
	this.bottomToolFrame.SetButtonEnable("play", false)
	this.bottomToolFrame.SetButtonEnable("stop", false)
	this.bottomToolFrame.SetSeek(0, 0, 0, "00:00/00:00")
	minVolume, maxVolume, curVolume := this.player.GetVolume()
	this.bottomToolFrame.SetVolume(minVolume, maxVolume, curVolume)
}

func (this *MainWindow) initBottomToolEmpty() {
	this.bottomToolFrame.SetButtonText("play", "暂停")
	this.bottomToolFrame.SetButtonEnable("play", true)
	this.bottomToolFrame.SetButtonEnable("stop", true)
	this.bottomToolFrame.SetSeek(0, 0, 0, "00:00/00:00")
	minVolume, maxVolume, curVolume := this.player.GetVolume()
	this.bottomToolFrame.SetVolume(minVolume, maxVolume, curVolume)
}
