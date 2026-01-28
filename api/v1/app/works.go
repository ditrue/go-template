package app

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/grocery/request"
	"github.com/gin-gonic/gin"
)

type WordsApi struct {
}

func (wordsApi *WordsApi) CreateWord(c *gin.Context) {
	var word app.Word
	err := c.ShouldBindJSON(&word)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = wordsService.CreateWord(c.Request.Context(), &word)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
}

func (wordsApi *WordsApi) CreateGrocery(c *gin.Context) {
	fmt.Println("====")
	req := request.GroceryCreate{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(req)

	err = wordsService.CreateGrocery(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData("返回数据", c)
}

func (wordsApi *WordsApi) GetGroceries(c *gin.Context) {
	req := request.GetGroceries{}
	// 查询数据库
	/*
			{
		  "code": 0,
		  "data": [
		    {
		      "ID": 1,
		      "CreatedAt": "2025-11-23T17:26:32.696608+08:00",
		      "UpdatedAt": "2025-11-23T17:26:32.696608+08:00",
		      "name": "3333",
		      "quantity": "1",
		      "category": "Spices"
		    }
		  ],
		  "msg": "成功"
		}
	*/
	if data, err := wordsService.GetGroceries(req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(data, c)
	}
}

func (wordsApi *WordsApi) GetWordsList(c *gin.Context) {
	if data, err := wordsService.GetWordsList(); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(data, c)
	}
}
