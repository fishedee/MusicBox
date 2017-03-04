package utils

import (
	. "github.com/fishedee/web"
	//"github.com/therecipe/qt/multimedia"
	"fmt"
	"github.com/therecipe/qt/widgets"
	"path"
)

type FileInfo struct {
	FilePath   string
	FileName   string
	FileFormat string
	Title      string
	Artist     string
	Duration   string
}

type FileDialog struct {
	Model
}

func NewFileDialog() *FileDialog {
	fileDialog := FileDialog{}
	InitModel(&fileDialog)
	fileDialog.init()
	return &fileDialog
}

func (this *FileDialog) init() {

}

func (this *FileDialog) loadMeta(files []string, index int, result []FileInfo, handler func([]FileInfo)) {
	if index == len(files) {
		handler(result)
		return
	}
	singleFile := files[index]
	if singleFile == "" {
		this.loadMeta(files, index+1, result, handler)
		return
	}
	loadMeta := false
	loadDuration := false
	fileInfo := FileInfo{}
	checkResult := func() {
		if loadMeta == false || loadDuration == false {
			return
		}
		result = append(result, fileInfo)
		this.loadMeta(files, index+1, result, handler)
	}
	fileInfo.FilePath = singleFile
	fileInfo.FileName = path.Base(singleFile)
	fileInfo.FileFormat = path.Ext(singleFile)
	player := NewPlayer()
	player.SetMetaListener(func() {
		metaData := player.GetMetaData()
		this.Log.Debug("meta %v", metaData)
		fileInfo.Title = metaData["title"]
		fileInfo.Artist = metaData["author"]
		loadMeta = true
		checkResult()
	})
	player.SetDurationChangeListener(func() {
		_, duration, _, _ := player.GetPosition()
		fileInfo.Duration = fmt.Sprintf("%02d:%02d", duration/60, duration%60)
		loadDuration = true
		checkResult()
	})
	player.SetFileName(singleFile)
}

func (this *FileDialog) Open(parent widgets.QWidget_ITF, handler func([]FileInfo)) []FileInfo {
	result := []FileInfo{}
	files := []string{"/Users/fishedee/Project/MusicBox/res/test.mp3"}
	//fileDialog := widgets.NewQFileDialog(parent, 0)
	//files := fileDialog.GetOpenFileNames(parent, "打开", "", "音频文件 (*.mp3 *.wma)", "", 0)
	this.loadMeta(files, 0, []FileInfo{}, handler)
	return result
}
