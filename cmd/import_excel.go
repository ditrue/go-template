package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/ditrue/go-template/global"
	"github.com/ditrue/go-template/model/app"
)

func importLetters() error {
	// A - Z a - z
	for i := 'A'; i <= 'Z'; i++ {
		letter := string(i)
		letterModel := app.Letter{
			Lowercase: strings.ToLower(letter),
			Uppercase: strings.ToUpper(letter),
		}
		global.GVA_DB.Create(&letterModel)
	}
	return nil
}

// 导入音频
func importAudio() error {
	// http save http://dict.youdao.com/dictvoice?type=0&audio=A
	var letters []app.Letter
	global.GVA_DB.Find(&letters)
	for _, letter := range letters {
		audioPath := fmt.Sprintf("http://dict.youdao.com/dictvoice?type=0&audio=%s", letter.Uppercase)
		resp, err := http.Get(audioPath)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		audioPath = fmt.Sprintf("audio/%s.mp3", letter.Uppercase)
		os.WriteFile(audioPath, body, 0644)
		letter.UppercaseAudioPath = audioPath
		global.GVA_DB.Save(&letter)

		// 小写字母
		audioPath = fmt.Sprintf("http://dict.youdao.com/dictvoice?type=0&audio=%s", letter.Lowercase)
		resp, err = http.Get(audioPath)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		audioPath = fmt.Sprintf("audio/%s.mp3", letter.Lowercase)
		os.WriteFile(audioPath, body, 0644)
		letter.LowercaseAudioPath = audioPath
		global.GVA_DB.Save(&letter)
	}
	return nil
}
