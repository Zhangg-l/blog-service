package dao

import "go_code/project8/blog-service/internal/model"

func (d *Dao) GetArticleTagByAId(articleId uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleId: articleId}
	return articleTag.GetByAId(d.engine)
}

func (d *Dao) GetArticleTagListByAIds(articleIds []uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{}
	return articleTag.ListByIds(d.engine, articleIds)
}

func (d *Dao) GetArticleTagByTId(tagId uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{TagId: tagId}
	return articleTag.ListByTId(d.engine)
}

func (d *Dao) CreateArticleTag(articleId, tagId uint32, createdBy string) error {
	articleTag := model.ArticleTag{TagId: tagId, ArticleId: articleId, Model: &model.Model{CreatedBy: createdBy}}
	return articleTag.Create(d.engine)
}

func (d *Dao) UpdateArticleTag(articleId, tagId uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{TagId: tagId, ArticleId: articleId, Model: &model.Model{ModifiedBy: modifiedBy}}
	value := map[string]interface{}{

		"article_id":  articleId,
		"tag_id":      tagId,
		"modified_by": modifiedBy,
	}
	return articleTag.UpdateOne(d.engine, value)
}

func (d *Dao) DeleteArticleTag(id uint32) error {
	articleTag := model.ArticleTag{ArticleId: id}

	return articleTag.Delete(d.engine)
}
