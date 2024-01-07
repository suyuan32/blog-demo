package publicapi

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"blog/internal/logic/publicapi"
	"blog/internal/svc"
	"blog/internal/types"
)

// swagger:route get /public/article/list publicapi GetPublicArticleList
//
// Get article public list | 获取Article公开列表
//
// Get article public list | 获取Article公开列表
//
// Responses:
//  200: ArticleListResp

func GetPublicArticleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticlePublicListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publicapi.NewGetPublicArticleListLogic(r.Context(), svcCtx)
		resp, err := l.GetPublicArticleList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
