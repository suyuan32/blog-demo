package category

import (
	"context"

	"blog/ent/category"
	"blog/ent/predicate"
	"blog/internal/svc"
	"blog/internal/types"
	"blog/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListLogic {
	return &GetCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryListLogic) GetCategoryList(req *types.CategoryListReq) (*types.CategoryListResp, error) {
	var predicates []predicate.Category
	if req.Title != nil {
		predicates = append(predicates, category.TitleContains(*req.Title))
	}
	if req.Remark != nil {
		predicates = append(predicates, category.RemarkContains(*req.Remark))
	}
	data, err := l.svcCtx.DB.Category.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.CategoryListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.CategoryInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        &v.ID,
					CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				},
				Title:  &v.Title,
				Remark: &v.Remark,
			})
	}

	return resp, nil
}
