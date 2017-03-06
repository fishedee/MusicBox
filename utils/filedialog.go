package utils

import (
	"fmt"
	"github.com/fishedee/web"
	"github.com/therecipe/qt/widgets"
	"io/ioutil"
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
	web.Model
}

func NewFileDialog() *FileDialog {
	fileDialog := FileDialog{}
	web.InitModel(&fileDialog)
	fileDialog.init()
	return &fileDialog
}

func (this *FileDialog) init() {

}

func (this *FileDialog) loadMeta(parent widgets.QWidget_ITF, files []string, index int, result []FileInfo, handler func([]FileInfo)) {
	if index == len(files) {
		handler(result)
		return
	}
	singleFile := files[index]
	if singleFile == "" {
		this.loadMeta(parent, files, index+1, result, handler)
		return
	}
	loadMeta := false
	loadDuration := false
	loadTimeout := false
	hasLoaded := false
	fileInfo := FileInfo{}
	checkResult := func() {
		if loadTimeout == false {
			if loadMeta == false || loadDuration == false {
				return
			}
		}
		if hasLoaded {
			return
		}
		result = append(result, fileInfo)
		this.loadMeta(parent, files, index+1, result, handler)
		hasLoaded = true
	}
	fileInfo.FilePath = singleFile
	fileInfo.FileName = path.Base(singleFile)
	fileInfo.FileFormat = path.Ext(singleFile)
	player := NewPlayer()
	player.SetErrorListener(func() {
		code, msg := player.GetError()
		this.Log.Debug("error", code, msg)
	})
	player.SetLoadedListener(func() {
		NewTimer().Sleep(parent, 500, func() {
			loadTimeout = true
			checkResult()
		})
	})
	player.SetMetaListener(func() {
		metaData := player.GetMetaData()
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

func (this *FileDialog) readDir(dirPath string) []string {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}
	}
	result := []string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		ext := path.Ext(name)
		if ext != ".mp3" && ext != ".wma" {
			continue
		}
		result = append(result, dirPath+"/"+name)
	}
	return result
}
func (this *FileDialog) Open(parent widgets.QWidget_ITF, handler func([]FileInfo)) []FileInfo {
	result := []FileInfo{}
	files := this.readDir("/Users/fishedee/Music/网易云音乐")
	//files := []string{"/Users/fishedee/Project/MusicBox/res/test.mp3"}
	//fileDialog := widgets.NewQFileDialog(parent, 0)
	//files := fileDialog.GetOpenFileNames(parent, "打开", "", "音频文件 (*.mp3 *.wma)", "", 0)
	this.loadMeta(parent, files, 0, []FileInfo{}, handler)
	return result
}
