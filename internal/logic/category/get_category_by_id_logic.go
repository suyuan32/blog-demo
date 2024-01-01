package category

import (
	"context"

	"blog/internal/svc"
	"blog/internal/types"
	"blog/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryByIdLogic {
	return &GetCategoryByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryByIdLogic) GetCategoryById(req *types.IDReq) (*types.CategoryInfoResp, error) {
	data, err := l.svcCtx.DB.Category.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.CategoryInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.CategoryInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        &data.ID,
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.UnixMilli()),
			},
			Title:  &data.Title,
			Remark: &data.Remark,
		},
	}, nil
}
