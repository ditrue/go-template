package app

import (
	"github.com/gin-gonic/gin"
)

type BookRouter struct{}

// InitBookRouter 初始化 书 路由信息
func (s *BookRouter) InitBookRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	boRouter := Router.Group("bo")
	boRouterWithoutRecord := Router.Group("bo")
	boRouterWithoutAuth := PublicRouter.Group("bo")
	{
		boRouter.POST("createBook", boApi.CreateBook)             // 新建书
		boRouter.DELETE("deleteBook", boApi.DeleteBook)           // 删除书
		boRouter.DELETE("deleteBookByIds", boApi.DeleteBookByIds) // 批量删除书
		boRouter.PUT("updateBook", boApi.UpdateBook)              // 更新书
		// 分享
	}
	{
		boRouterWithoutRecord.GET("findBook", boApi.FindBook) // 根据ID获取书
	}
	{
		boRouterWithoutAuth.GET("getBookPublic", boApi.GetBookPublic) // 书开放接口
		boRouterWithoutAuth.GET(":id/shareBook", boApi.ShareBook)
		boRouterWithoutAuth.GET("shareBooks", boApi.ShareBooks)
		boRouterWithoutAuth.GET("getBookList", boApi.GetBookList) // 获取书列表
	}
}
