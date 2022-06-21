package service

import (
	"go_code/project8/blog-service/internal/dao"
	"go_code/project8/blog-service/internal/model"
	"go_code/project8/blog-service/pkg/app"
)

type ArticleRequest struct {
	Id    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	TagId uint32 `form:"tag_id" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	TagId         uint32 `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	Content       string `form:"content" binding:"required,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	CreatedBy     string `form:"created_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	Id            uint32 `form:"id" binding:"required,gte=1"`
	TagId         uint32 `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"min=2,max=100"`
	Desc          string `form:"desc" binding:"min=2,max=255"`
	Content       string `form:"content" binding:"min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"url"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	Id uint32 `form:"id" binding:"required,gte=1"`
}

type Article struct {
	Id            uint32     `json:"id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Content       string     `json:"content"`
	CoverImageUrl string     `json:"cover_image_url"`
	State         uint8      `json:"state"`
	Tag           *model.Tag `json:"tag"`
}

func (svc *Service) GetArticle(param *ArticleRequest) (*Article, error) {
	article, err := svc.dao.GetArticle(param.Id, param.State)
	if err != nil {
		return nil, err
	}

	articleTag, err := svc.dao.GetArticleTagByAId(article.Id)
	if err != nil {
		return nil, err
	}

	tag, err := svc.dao.GetTag(articleTag.TagId, model.STATE_OPEN)
	if err != nil {
		return nil, err
	}

	return &Article{
		Id:            article.Id,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State:         article.State,
		Tag:           tag,
	}, nil
}

func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*Article, int, error) {
	articleCount, err := svc.dao.CountArticleListByTagId(param.TagId, param.State)
	if err != nil {
		return nil, 0, err
	}

	articles, err := svc.dao.GetArticleListByTagId(param.TagId, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var articleList []*Article
	for _, article := range articles {
		articleList = append(articleList, &Article{
			Id:            article.ArticleId,
			Title:         article.ArticleTitle,
			Desc:          article.ArticleDesc,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
			Tag:           &model.Tag{Model: &model.Model{Id: article.TagId}, Name: article.TagName},
		})
	}

	return articleList, articleCount, nil
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	article, err := svc.dao.CreateArticle(&dao.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		CreatedBy:     param.CreatedBy,
	})
	if err != nil {
		return err
	}

	err = svc.dao.CreateArticleTag(article.Id, param.TagId, param.CreatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	err := svc.dao.UpdateArticle(&dao.Article{
		Id:            param.Id,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		ModifiedBy:    param.ModifiedBy,
	})
	if err != nil {
		return err
	}

	err = svc.dao.UpdateArticleTag(param.Id, param.TagId, param.ModifiedBy)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	err := svc.dao.DeleteArticle(param.Id)
	if err != nil {
		return err
	}

	err = svc.dao.DeleteArticleTag(param.Id)
	if err != nil {
		return err
	}

	return nil
}
