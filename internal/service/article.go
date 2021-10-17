package service

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
	"time"
)
type ArticleCountRequest struct {
	Title string `form:"title" `
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}
type ArticleListRequest struct {
	Title string `form:"title"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}
type ArticleContRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
type CreateArticleRequest struct {
	Title string `form:"title" binding:"required,min=1"`
	Desc string `form:"desc"`
	Content string `form:"content" binding:"required,min=1"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}
type UpdateArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
	Title string `form:"title" `
	Desc string `form:"desc"`
	Content string `form:"content""`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}
type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}


func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title,param.Desc,param.Content,param.State,param.CreatedBy)
}
func (svc *Service) GetArticle(param *ArticleContRequest) (model.Article,error) {
	return svc.dao.GetArticle(param.ID)
}


func (svc *Service) GetArticleList(param *ArticleListRequest,pager *app.Pager) ([]*model.ArticleList,int,error) {
	articledCount,err := svc.dao.CountArticle(param.Title,param.State)
	if err != nil {
		return nil, 0, err
	}
	articleds,err :=  svc.dao.GetArticleList(param.Title,param.State,pager.Page,pager.PageSize)
	if err != nil {
		return nil, 0, err
	}
	var articledList []*model.ArticleList
	var id uint32 = 1
	for _,articled := range articleds{
		articledList = append(articledList,&model.ArticleList{
			ID:         id,
			Title:      articled.Title,
			Desc:       articled.Desc,
			CreatedBy:  articled.CreatedBy,
			CreatedOn:  time.Unix(int64(articled.CreatedOn), 0).Format("2006-01-02 15:04:05"),
			State:      articled.State,
			ModifiedBy: articled.ModifiedBy,
			ModifiedOn: time.Unix(int64(articled.ModifiedOn), 0).Format("2006-01-02 15:04:05"),
		})
		id +=1
	}
	return articledList,articledCount,nil
}
func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(param.ID,param.Title,param.Desc,param.Content,param.State,param.ModifiedBy)
}
func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.ID)
}