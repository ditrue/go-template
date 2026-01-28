package app

import (
	"github.com/gin-gonic/gin"
)

type WordsRouter struct{}

// InitWordsRouter 初始化 单词 路由信息
func (s *WordsRouter) InitWordsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	wordsRouter := Router.Group("words")
	{
		// 获取words列表
		wordsRouter.GET("", wordsApi.GetWordsList)
		// test
		wordsRouter.GET("test", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Hello, World!"})
		})
		wordsRouter.POST("createWord", wordsApi.CreateWord) // 新建单词
		// 测试接口 groceries
		wordsRouter.POST("groceries", wordsApi.CreateGrocery)
		wordsRouter.GET("groceries", wordsApi.GetGroceries)
	}
}
