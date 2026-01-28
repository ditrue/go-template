package app

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	"github.com/flipped-aurora/gin-vue-admin/server/model/grocery/request"
)

type WordsService struct{}

func (wordsService *WordsService) CreateWord(ctx context.Context, word *app.Word) (err error) {
	err = global.GVA_DB.Create(word).Error
	return err
}

func (wordsService *WordsService) CreateGrocery(req request.GroceryCreate) (err error) {
	grocery := app.Grocery{
		Name:     req.Name,
		Quantity: req.Quantity,
		Category: req.Category,
	}
	err = global.GVA_DB.Create(&grocery).Error
	if err != nil {
		return err
	}
	return nil
}

func (wordsService *WordsService) GetGroceries(req request.GetGroceries) (data []app.Grocery, err error) {
	err = global.GVA_DB.Model(&app.Grocery{}).Where("name LIKE ?", "%"+req.Keyword+"%").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetWordsList 获取words列表
func (wordsService *WordsService) GetWordsList() (data []app.Word, err error) {
	err = global.GVA_DB.Model(&app.Word{}).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
