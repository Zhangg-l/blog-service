package dao

import (
	"go_code/project8/blog-service/internal/model"
	"go_code/project8/blog-service/pkg/app"
)

type Article struct {
	Id            uint32 `json:"id"`
	TagId         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

func (d *Dao) UpdateArticle(param *Article) error {

	article := model.Article{Model: &model.Model{Id: param.Id}}

	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}
	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.Title
	}

	if param.Desc != "" {
		values["Desc"] = param.Desc
	}

	if param.Content != "" {
		values["content"] = param.Content
	}

	return article.Update(d.engine, values)
}

func (d *Dao) CreateArticle(param *Article) (*model.Article, error) {

	article := model.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Model:         &model.Model{CreatedBy: param.CreatedBy},
	}
	return article.Create(d.engine)
}

func (d *Dao) GetArticle(id uint32, state uint8) (model.Article, error) {

	article := model.Article{
		Model: &model.Model{Id: id},
		State: state,
	}
	return article.Get(d.engine)
}

func (d *Dao) DeleteArticle(id uint32) error {

	article := model.Article{
		Model: &model.Model{Id: id},
	}
	return article.Delete(d.engine)
}

func (d *Dao) CountArticleListByTagId(tid uint32, state uint8) (int, error) {

	article := model.Article{
		State: state,
	}
	return article.CountByTagId(d.engine, tid)
}

func (d *Dao) GetArticleListByTagId(tid uint32, state uint8, page, pageSize int) ([]*model.ArticleRow, error) {

	article := model.Article{
		State: state,
	}
	return article.ListByTagId(d.engine, tid, app.GetPageOffset(page, pageSize), pageSize)
}
