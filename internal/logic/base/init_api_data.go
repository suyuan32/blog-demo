package base

import (
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

func (l *InitDatabaseLogic) insertApiData() (err error) {
	// Category

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Blog"),
		Path:        pointy.GetPointer("/category/create"),
		Description: pointy.GetPointer("apiDesc.createCategory"),
		ApiGroup:    pointy.GetPointer("category"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Blog"),
		Path:        pointy.GetPointer("/category/update"),
		Description: pointy.GetPointer("apiDesc.updateCategory"),
		ApiGroup:    pointy.GetPointer("category"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Blog"),
		Path:        pointy.GetPointer("/category/delete"),
		Description: pointy.GetPointer("apiDesc.deleteCategory"),
		ApiGroup:    pointy.GetPointer("category"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Blog"),
		Path:        pointy.GetPointer("/category/list"),
		Description: pointy.GetPointer("apiDesc.getCategoryList"),
		ApiGroup:    pointy.GetPointer("category"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Blog"),
		Path:        pointy.GetPointer("/category"),
		Description: pointy.GetPointer("apiDesc.getCategoryById"),
		ApiGroup:    pointy.GetPointer("category"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	// Article

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Blog"),
		Path:        pointy.GetPointer("/article/create"),
		Description: pointy.GetPointer("apiDesc.createArticle"),
		ApiGroup:    pointy.GetPointer("article"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Blog"),
		Path:        pointy.GetPointer("/article/update"),
		Description: pointy.GetPointer("apiDesc.updateArticle"),
		ApiGroup:    pointy.GetPointer("article"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Blog"),
		Path:        pointy.GetPointer("/article/delete"),
		Description: pointy.GetPointer("apiDesc.deleteArticle"),
		ApiGroup:    pointy.GetPointer("article"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Blog"),
		Path:        pointy.GetPointer("/article/list"),
		Description: pointy.GetPointer("apiDesc.getArticleList"),
		ApiGroup:    pointy.GetPointer("article"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Blog"),
		Path:        pointy.GetPointer("/article"),
		Description: pointy.GetPointer("apiDesc.getArticleById"),
		ApiGroup:    pointy.GetPointer("article"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	return err
}
