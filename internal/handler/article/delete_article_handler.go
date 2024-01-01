package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"blog/internal/logic/article"
	"blog/internal/svc"
	"blog/internal/types"
)

// swagger:route post /article/delete article DeleteArticle
//
// Delete article information | 删除Article信息
//
// Delete article information | 删除Article信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UUIDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUIDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewDeleteArticleLogic(r.Context(), svcCtx)
		resp, err := l.DeleteArticle(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
