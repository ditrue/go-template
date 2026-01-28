package initialize

import (
	"github.com/ditrue/go-template/global"
	"github.com/ditrue/go-template/model/app"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(
		app.Book{},
		app.Share{},
		app.Demo{},
		app.Word{},
		app.Letter{},
		app.GroupPhonic{},
		app.Grocery{},
		app.User{},
	)
	if err != nil {
		return err
	}
	return nil
}
