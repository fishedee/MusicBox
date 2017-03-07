package models

import (
	"bufio"
	"fmt"
	. "github.com/fishedee/encoding"
	. "github.com/fishedee/language"
	. "github.com/fishedee/util"
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/core"
	"io"
	"strconv"
	"strings"
)

type lrc struct {
	time int
	lrc  string
}

type lrcData struct {
	err  string
	data []lrc
}

type MusicLrc struct {
	Model
	lrcList          map[string]lrcData
	downloadListener func()
}

func NewMusicLrc() *MusicLrc {
	musicLrc := MusicLrc{}
	InitModel(&musicLrc)
	musicLrc.init()
	return &musicLrc
}

func (this *MusicLrc) init() {
	this.lrcList = map[string]lrcData{}
}

func (this *MusicLrc) DownloadLrc(parent core.QObject_ITF, music Music) {
	var resultErr error
	var resultData []lrc
	timer := core.NewQTimer(parent)
	timer.SetSingleShot(true)
	timer.ConnectTimeout(func() {
		this.finish(music, resultErr, resultData)
	})

	go func() {
		defer CatchCrash(func(e Exception) {
			this.Log.Debug("msg:%v,stack:%v", e.GetMessage(), e.GetStackTrace())
			resultErr = e
			resultData = nil
			timer.Start(0)
		})
		keyword := ""
		if music.Title == "" {
			return
		}
		keyword = music.Title
		if music.Artist != "" {
			keyword += "|" + music.Artist
		}
		var durationMinute, durationSecond int
		_, err := fmt.Sscanf(music.Duration, "%d:%d", &durationMinute, &durationSecond)
		if err != nil {
			panic(err)
		}
		durationMilliSecond := (durationMinute*60 + durationSecond) * 1000

		var result map[string]interface{}

		err = DefaultAjaxPool.Get(&Ajax{
			Url: "http://lyrics.kugou.com/search",
			Data: map[string]interface{}{
				"ver":      1,
				"man":      "yes",
				"client":   "pc",
				"keyword":  keyword,
				"duration": durationMilliSecond,
			},
			ResponseData:     &result,
			ResponseDataType: "json",
		})
		if err != nil {
			panic(err)
		}
		status := int(result["status"].(float64))
		if status != 200 {
			if status == 404 {
				panic("找不到歌词")
			} else {
				panic("歌词错误码为" + strconv.Itoa(status))
			}
		}
		candidates := result["candidates"].([]interface{})
		firstCandidate := candidates[0].(map[string]interface{})

		err = DefaultAjaxPool.Get(&Ajax{
			Url: "http://lyrics.kugou.com/download",
			Data: map[string]interface{}{
				"ver":       1,
				"client":    "pc",
				"id":        firstCandidate["id"].(string),
				"accesskey": firstCandidate["accesskey"].(string),
				"fmt":       "lrc",
				"charset":   "utf8",
			},
			ResponseData:     &result,
			ResponseDataType: "json",
		})
		if err != nil {
			panic(err)
		}
		status = int(result["status"].(float64))
		if status != 200 {
			panic("下载歌词错误码为" + strconv.Itoa(status))
		}
		content64 := result["content"].(string)
		content, err := DecodeBase64(content64)
		if err != nil {
			panic(err)
		}

		contentReader := strings.NewReader(string(content))
		contentBufReader := bufio.NewReader(contentReader)
		resultLrc := []lrc{}
		for {
			contentLine, err := contentBufReader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				} else {
					panic(err)
				}
			}
			var miniute, second, miliSecond int
			var lineData string
			_, err = fmt.Sscanf(string(contentLine), "[%d:%d.%d]", &miniute, &second, &miliSecond)
			if err != nil {
				panic(err)
			}
			lineData = contentLine[10 : len(contentLine)-2]
			resultLrc = append(resultLrc, lrc{
				time: miniute*60 + second,
				lrc:  lineData,
			})
		}
		resultErr = nil
		resultData = resultLrc
		timer.Start(0)
	}()
}

func (this *MusicLrc) finish(music Music, err error, result []lrc) {
	if err != nil {
		this.lrcList[music.FilePath] = lrcData{
			err:  err.Error(),
			data: nil,
		}
	} else {
		this.lrcList[music.FilePath] = lrcData{
			err:  "",
			data: result,
		}
	}
	if this.downloadListener != nil {
		this.downloadListener()
	}
}

func (this *MusicLrc) SetDownloadedListener(listener func()) {
	this.downloadListener = listener
}

func (this *MusicLrc) HasLrc(music Music) bool {
	_, isOk := this.lrcList[music.FilePath]
	return isOk
}

func (this *MusicLrc) GetLrcStatus(music Music) string {
	result, isOk := this.lrcList[music.FilePath]
	if isOk == false {
		return ""
	}
	return result.err
}

func (this *MusicLrc) GetLrcData(music Music) []string {
	result, isOk := this.lrcList[music.FilePath]
	if isOk == false {
		return nil
	}
	resultStr := []string{}
	for _, singleLine := range result.data {
		resultStr = append(resultStr, singleLine.lrc)
	}
	return resultStr
}

func (this *MusicLrc) GetLrcProgress(music Music, progress int) int {
	result, isOk := this.lrcList[music.FilePath]
	if isOk == false {
		return -1
	}
	for i := len(result.data) - 1; i >= 0; i-- {
		if result.data[i].time < progress {
			return i
		}
	}
	return 0
}
