package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"blog/internal/logic/article"
	"blog/internal/svc"
	"blog/internal/types"
)

// swagger:route post /article/create article CreateArticle
//
// Create article information | 创建Article
//
// Create article information | 创建Article
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ArticleInfo
//
// Responses:
//  200: BaseMsgResp

func CreateArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewCreateArticleLogic(r.Context(), svcCtx)
		resp, err := l.CreateArticle(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
