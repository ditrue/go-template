package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ditrue/go-template/global"
	"github.com/ditrue/go-template/model/app"
)

func importWord() error {
	/*
		leg
		step
		these
		evening
		Chinese
		theme
		fruit
		banana
		orange
		grape
		watermelon
		strawberry
	*/
	words := []string{
		"leg",
		"step",
		"these",
		"evening",
		"Chinese",
		"theme",
		"fruit",
		"banana",
		"orange",
		"grape",
		"watermelon",
		"strawberry",
	}
	for _, word := range words {
		wordModel := app.Word{
			Word: word,
		}
		global.GVA_DB.Create(&wordModel)
	}
	return nil
}

// 导入音频
func importWordAudio() error {
	var words []app.Word
	global.GVA_DB.Find(&words)
	for _, word := range words {
		audioPath := fmt.Sprintf("http://dict.youdao.com/dictvoice?type=0&audio=%s", word.Word)
		resp, err := http.Get(audioPath)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		audioPath = fmt.Sprintf("audio/%s.mp3", word.Word)
		os.WriteFile(audioPath, body, 0644)
		word.AudioPath = audioPath
		global.GVA_DB.Save(&word)
	}
	return nil
}
