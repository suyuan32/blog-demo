package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"blog/internal/logic/category"
	"blog/internal/svc"
	"blog/internal/types"
)

// swagger:route post /category/update category UpdateCategory
//
// Update category information | 更新Category
//
// Update category information | 更新Category
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CategoryInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewUpdateCategoryLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCategory(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
