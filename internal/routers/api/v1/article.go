package v1

import (
	"blog-service/global"
	"blog-service/internal/service"
	"blog-service/pkg/app"
	"blog-service/pkg/convert"
	"blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

// @Summary 文章详情
// @Produce  json
// @Param id path int true "文章ID"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article/{id} [get]
func (a Article) Get(c *gin.Context) {
	param := service.ArticleContRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v",errs)
		response.ToErrorResponse(errcode.ErrorGetArticleContFail)
		return
	}
	svc := service.New(c.Request.Context())
	article, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteArticle err %v",err)
		response.ToErrorResponse(errcode.ErrorGetArticleContFail)
		return
	}
	response.ToResponse(article)
	return
}
// @Summary 获取文章列表
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article [get]
func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs %v",errs)
		response.ToErrorResponse(errcode.ErrorGetArticleListFail)
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c),PageSize: app.GetPageSize(c)}
	articles,pageRows,err:=svc.GetArticleList(&param,&pager)
	if err != nil {
		global.Logger.Errorf("svc.GetArticleList err %v",err)
		response.ToErrorResponse(errcode.ErrorGetArticleListFail)
		return
	}
	response.ToResponseList(articles,pageRows)
	return

}
// @Summary 新增文章
// @Produce  json
// @Param title body string true "文章标题"
// @Param desc body string true "文章简述"
// @Param content body string true "文章内容"
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者"
// @Success 200 {object} errcode.Error "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article [post]
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid,errs:= app.BindAndValid(c,&param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs %v",errs)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&service.CreateArticleRequest{
		Title:     param.Title,
		Desc:      param.Desc,
		Content:   param.Content,
		CreatedBy: param.CreatedBy,
		State:     param.State,
	})
	if err != nil {
		global.Logger.Errorf("svc.CreateArticle err %v",err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	response.ToErrorResponse(errcode.Success)
	return
}
// @Summary 更新文章
// @Produce  json
// @Param id path int true "文章ID"
// @Param title body string true "文章标题"
// @Param desc body string true "文章简述"
// @Param content body string true "文章内容"
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者"
// @Success 200 {object} errcode.Error "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article/{id} [put]
func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v",errs)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateArticle err %v",err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	response.ToErrorResponse(errcode.Success)
	return
}
// @Summary 删除文章
// @Produce  json
// @Param id path int true "文章ID"
// @Success 200 {object} errcode.Error "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v",errs)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteArticle err %v",err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	response.ToErrorResponse(errcode.Success)
	return
}
