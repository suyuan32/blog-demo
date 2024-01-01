package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"blog/internal/logic/category"
	"blog/internal/svc"
	"blog/internal/types"
)

// swagger:route post /category/delete category DeleteCategory
//
// Delete category information | 删除Category信息
//
// Delete category information | 删除Category信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewDeleteCategoryLogic(r.Context(), svcCtx)
		resp, err := l.DeleteCategory(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
