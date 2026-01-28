package app

import api "github.com/ditrue/go-template/api/v1"

type RouterGroup struct {
	BookRouter
	WordsRouter
	AuthRouter
}

var boApi = api.ApiGroupApp.AppApiGroup.BookApi
var wordsApi = api.ApiGroupApp.AppApiGroup.WordsApi
