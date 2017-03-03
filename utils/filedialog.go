package utils

import (
	. "github.com/fishedee/web"
	//"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"
	"path"
)

type FileInfo struct {
	FilePath   string
	FileName   string
	FileFormat string
	Title      string
	Artist     string
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

func (this *FileDialog) Open(parent widgets.QWidget_ITF) []FileInfo {
	this.Log.Debug("%v", "ii")
	result := []FileInfo{}
	fileDialog := widgets.NewQFileDialog(parent, 0)
	files := fileDialog.GetOpenFileNames(parent, "打开", "", "音频文件 (*.mp3 *.wma)", "", 0)
	this.Log.Debug("%v", files)
	for _, singleFile := range files {
		player := NewPlayer()
		player.SetMetaListener(func() {
			this.Log.Debug("meta %v", player.GetMetaData())
		})
		player.SetFileName(singleFile)
		title := ""
		artist := ""
		name := path.Base(singleFile)
		format := path.Ext(singleFile)
		result = append(result, FileInfo{
			singleFile,
			name,
			format,
			title,
			artist,
		})
	}
	this.Log.Debug("%v", result)
	return result
}
