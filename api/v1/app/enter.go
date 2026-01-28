package app

import "github.com/ditrue/go-template/service"

type ApiGroup struct {
	BookApi
	WordsApi
	LoginApi
}

var boService = service.ServiceGroupApp.AppServiceGroup.BookService
var wordsService = service.ServiceGroupApp.AppServiceGroup.WordsService
var usersService = service.ServiceGroupApp.AppServiceGroup.UsersService
