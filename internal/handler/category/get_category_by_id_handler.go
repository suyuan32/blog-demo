package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"blog/internal/logic/category"
	"blog/internal/svc"
	"blog/internal/types"
)

// swagger:route post /category category GetCategoryById
//
// Get category by ID | 通过ID获取Category
//
// Get category by ID | 通过ID获取Category
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: CategoryInfoResp

func GetCategoryByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewGetCategoryByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetCategoryById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
