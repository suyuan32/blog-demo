package article

import (
	"context"

	"blog/internal/svc"
	"blog/internal/types"
	"blog/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByIdLogic {
	return &GetArticleByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleByIdLogic) GetArticleById(req *types.UUIDReq) (*types.ArticleInfoResp, error) {
	data, err := l.svcCtx.DB.Article.Get(l.ctx, uuidx.ParseUUIDString(req.Id))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.ArticleInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.ArticleInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        pointy.GetPointer(data.ID.String()),
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.UnixMilli()),
			},
			Title:   &data.Title,
			Content: &data.Content,
			Keyword: &data.Keyword,
			Visit:   &data.Visit,
		},
	}, nil
}
