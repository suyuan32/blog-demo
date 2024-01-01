package article

import (
	"context"

	"blog/ent/article"
	"blog/ent/predicate"
	"blog/internal/svc"
	"blog/internal/types"
	"blog/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListLogic {
	return &GetArticleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleListLogic) GetArticleList(req *types.ArticleListReq) (*types.ArticleListResp, error) {
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
	data, err := l.svcCtx.DB.Article.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.ArticleListResp{}
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
				Title:   &v.Title,
				Content: &v.Content,
				Keyword: &v.Keyword,
				Visit:   &v.Visit,
			})
	}

	return resp, nil
}
