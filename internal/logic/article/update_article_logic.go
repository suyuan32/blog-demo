package article

import (
	"context"

	"blog/internal/svc"
	"blog/internal/types"
	"blog/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateArticleLogic) UpdateArticle(req *types.ArticleInfo) (*types.BaseMsgResp, error) {
	query := l.svcCtx.DB.Article.UpdateOneID(uuidx.ParseUUIDString(*req.Id)).
		SetNotNilTitle(req.Title).
		SetNotNilContent(req.Content).
		SetNotNilKeyword(req.Keyword).
		SetNotNilVisit(req.Visit).
		SetNotNilStatus(req.Status)

	if req.CategoryId != nil {
		query.SetCategoryID(*req.CategoryId)
	}

	err := query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
