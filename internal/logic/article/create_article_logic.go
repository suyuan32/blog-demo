package article

import (
	"context"

	"blog/internal/svc"
	"blog/internal/types"
	"blog/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateArticleLogic) CreateArticle(req *types.ArticleInfo) (*types.BaseMsgResp, error) {
	_, err := l.svcCtx.DB.Article.Create().
		SetNotNilTitle(req.Title).
		SetNotNilContent(req.Content).
		SetNotNilKeyword(req.Keyword).
		SetNotNilVisit(req.Visit).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
