package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"blog/internal/logic/category"
	"blog/internal/svc"
	"blog/internal/types"
)

// swagger:route post /category/list category GetCategoryList
//
// Get category list | 获取Category列表
//
// Get category list | 获取Category列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CategoryListReq
//
// Responses:
//  200: CategoryListResp

func GetCategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewGetCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.GetCategoryList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
