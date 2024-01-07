package publicapi

import (
	"blog/ent/article"
	"blog/internal/utils/dberrorhandler"
	"context"
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublicArticleByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublicArticleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicArticleByIdLogic {
	return &GetPublicArticleByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPublicArticleByIdLogic) GetPublicArticleById(req *types.UUIDPathReq) (resp *types.ArticleInfoResp, err error) {
	data, err := l.svcCtx.DB.Article.Query().Where(article.IDEQ(uuidx.ParseUUIDString(req.Id)),
		article.StatusEQ(common.StatusNormal)).WithCategory().Only(l.ctx)
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
			Title:      &data.Title,
			Content:    &data.Content,
			Keyword:    &data.Keyword,
			Visit:      &data.Visit,
			CategoryId: &data.Edges.Category.ID,
			Status:     &data.Status,
		},
	}, nil
}
