package publicapi

import (
	"blog/ent/article"
	"blog/ent/predicate"
	"blog/internal/utils/dberrorhandler"
	"context"
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublicArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublicArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicArticleListLogic {
	return &GetPublicArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPublicArticleListLogic) GetPublicArticleList(req *types.ArticlePublicListReq) (resp *types.ArticleListResp, err error) {
	var predicates []predicate.Article
	if req.Title != nil {
		predicates = append(predicates, article.TitleContains(*req.Title))
	}
	if req.Content != nil {
		predicates = append(predicates, article.ContentContains(*req.Content))
	}
	if req.Keyword != nil {
		predicates = append(predicates, article.KeywordContains(*req.Keyword))
	}

	predicates = append(predicates, article.StatusEQ(common.StatusNormal))

	data, err := l.svcCtx.DB.Article.Query().Where(predicates...).WithCategory().Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp = &types.ArticleListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.ArticleInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        pointy.GetPointer(v.ID.String()),
					CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				},
				Title:      &v.Title,
				Content:    &v.Content,
				Keyword:    &v.Keyword,
				Visit:      &v.Visit,
				CategoryId: &v.Edges.Category.ID,
				Status:     &v.Status,
			})
	}

	return resp, nil
}
