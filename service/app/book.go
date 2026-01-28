package app

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/response"
)

type BookService struct{}

// CreateBook 创建书记录
// Author [yourname](https://github.com/yourname)
func (boService *BookService) CreateBook(ctx context.Context, bo *app.Book) (err error) {
	err = global.GVA_DB.Create(bo).Error
	return err
}

// DeleteBook 删除书记录
// Author [yourname](https://github.com/yourname)
func (boService *BookService) DeleteBook(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&app.Book{}, "id = ?", ID).Error
	return err
}

// DeleteBookByIds 批量删除书记录
// Author [yourname](https://github.com/yourname)
func (boService *BookService) DeleteBookByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]app.Book{}, "id in ?", IDs).Error
	return err
}

// UpdateBook 更新书记录
// Author [yourname](https://github.com/yourname)
func (boService *BookService) UpdateBook(ctx context.Context, bo app.Book) (err error) {
	err = global.GVA_DB.Model(&app.Book{}).Where("id = ?", bo.ID).Updates(&bo).Error
	return err
}

// GetBook 根据ID获取书记录
// Author [yourname](https://github.com/yourname)
func (boService *BookService) GetBook(ctx context.Context, ID string) (bo app.Book, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&bo).Error
	return
}

// GetBookInfoList 分页获取书记录
// Author [yourname](https://github.com/yourname)
func (boService *BookService) GetBookInfoList(ctx context.Context, info appReq.BookSearch) (list []app.Book, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.Book{})
	var bos []app.Book
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&bos).Error
	return bos, total, err
}
func (boService *BookService) GetBookPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// ShareBook 分享书
func (boService *BookService) ShareBook(req *appReq.CreateBook) (err error) {
	fmt.Println(req.ID, req.ToUserID)
	var share app.Share
	share.BookID = req.ID
	share.ToUserID = req.ToUserID
	share.UserID = 1
	err = global.GVA_DB.Create(&share).Error
	return err
}

// GetShareBooks 获取分享书
func (boService *BookService) GetShareBooks() (books []response.BookResponse, err error) {
	var resp []app.Share
	fmt.Println(123)
	err = global.GVA_DB.Debug().Preload("Book").Find(&resp).Error
	fmt.Println(resp)
	for _, share := range resp {
		books = append(books, response.BookResponse{
			ID:   share.BookID,
			Book: *share.Book.Book,
		})
	}
	fmt.Println(books)
	return books, err
}
