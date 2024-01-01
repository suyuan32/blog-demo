package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"blog/internal/logic/article"
	"blog/internal/svc"
	"blog/internal/types"
)

// swagger:route post /article/update article UpdateArticle
//
// Update article information | 更新Article
//
// Update article information | 更新Article
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ArticleInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewUpdateArticleLogic(r.Context(), svcCtx)
		resp, err := l.UpdateArticle(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
