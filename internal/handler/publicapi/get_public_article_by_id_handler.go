package publicapi

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"blog/internal/logic/publicapi"
	"blog/internal/svc"
	"blog/internal/types"
)

// swagger:route get /public/article/{id} publicapi GetPublicArticleById
//
// Get article by ID | 通过ID获取Article
//
// Get article by ID | 通过ID获取Article
//
// Responses:
//  200: ArticleInfoResp

func GetPublicArticleByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUIDPathReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publicapi.NewGetPublicArticleByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetPublicArticleById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
