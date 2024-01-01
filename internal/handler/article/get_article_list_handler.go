package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"blog/internal/logic/article"
	"blog/internal/svc"
	"blog/internal/types"
)

// swagger:route post /article/list article GetArticleList
//
// Get article list | 获取Article列表
//
// Get article list | 获取Article列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ArticleListReq
//
// Responses:
//  200: ArticleListResp

func GetArticleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewGetArticleListLogic(r.Context(), svcCtx)
		resp, err := l.GetArticleList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
