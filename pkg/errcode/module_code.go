package errcode

var (
	ErrorGetTagListFail     = NewError(200101, "获取标签列表失败")
	ErrorCreateTagFail      = NewError(200102, "创建标签失败")
	ErrorUpdateTagFail      = NewError(200103, "更新标签失败")
	ErrorDeleteTagFail      = NewError(200104, "删除标签失败")
	ErrorCountTagFail       = NewError(200105, "统计标签失败")
	ErrorGetArticleListFail = NewError(200200, "获取文章列表失败")
	ErrorGetArticleContFail = NewError(200201, "获取文章详情失败")
	ErrorCreateArticleFail  = NewError(200202, "创建文章失败")
	ErrorUpdateArticleFail  = NewError(200203, "更新文章失败")
	ErrorDeleteArticleFail  = NewError(200204, "删除文章失败")
	ErrorCountArticleFail   = NewError(200205, "统计文章失败")
)
