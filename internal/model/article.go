package model

import (
	"blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	State   uint8  `json:"state"`
}

type ArticleList struct {
	ID         uint32 `json:"id"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	CreatedBy  string `json:"created_by"`
	CreatedOn  string `json:"created_on"`
	State      uint8  `json:"state"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn string `json:"modified_on"`
}
type ArticleSwagger struct {
	List  []*ArticleList
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}
func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	err := db.Where("id =? and is_del=?",a.ID,0).First(&article).Error
	if err !=nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	return article,nil
}
func (a Article) Count(db *gorm.DB) (int, error) {
	var count int
	if a.Title != "" {
		db = db.Where("title =?", a.Title)
	}
	db = db.Where("state =?", a.State)
	err := db.Model(&a).Where("is_del =?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var artcle []*Article
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if a.Title != "" {
		db = db.Where("title =?", a.Title)
	}
	db = db.Where("state =?", a.State)
	err = db.Where("is_del =?", 0).Find(&artcle).Error
	if err != nil {
		return nil, err
	}
	return artcle, nil
}
func (a Article) Update(db *gorm.DB, value interface{}) error {
	return db.Model(&Article{}).Where("id=? and is_del=?", a.ID, 0).Update(value).Error
}
func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? and is_del =?",a.ID,0).Delete(&a).Error
}