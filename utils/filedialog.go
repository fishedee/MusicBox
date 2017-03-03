package utils

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/multimedia"
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
	result := []FileInfo{}
	fileDialog := widgets.NewQFileDialog(parent, 0)
	files := fileDialog.GetOpenFileNames(parent, "打开", "", "音频文件 (*.mp3 *.wma)", "", 0)
	for _, singleFile := range files {
		decoder := multimedia.NewQAudioDecoder(parent)
		decoder.SetSourceFilename(singleFile)
		title := decoder.MetaData("title")
		artist := decoder.MetaData("artist")
		name := path.Base(singleFile)
		format := path.Ext(singleFile)
		result = append(result, FileInfo{
			singleFile,
			name,
			format,
			title.ToString(),
			artist.ToString(),
		})
	}
	return result
}
