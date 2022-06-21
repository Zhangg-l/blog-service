package model

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	*Model
	TagId     uint32 `json:"tag_id"`
	ArticleId uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (a ArticleTag) GetByAId(db *gorm.DB) (ArticleTag, error) {

	var articleTag ArticleTag
	err := db.Where("article_id=? AND is_del = ?", a.ArticleId, 0).First(&articleTag).Error

	if err != nil {
		return articleTag, err
	}
	return articleTag, nil
}

func (a ArticleTag) ListByTId(db *gorm.DB) (ArticleTag, error) {

	var articleTag ArticleTag
	err := db.Where("tag_id=? AND is_del = ?", a.TagId, 0).First(&articleTag).Error

	if err != nil {
		return articleTag, err
	}
	return articleTag, nil
}

func (a ArticleTag) ListByIds(db *gorm.DB, ids []uint32) ([]*ArticleTag, error) {

	var articleTags []*ArticleTag

	err := db.Where("article_id IN (?) AND is_del = ?", ids, 0).Find(&articleTags).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articleTags, nil
}

func (a ArticleTag) Create(db *gorm.DB) error {

	if err := db.Create(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a ArticleTag) UpdateOne(db *gorm.DB, value interface{}) error {

	if err := db.Model(&a).Where("article_id=? AND is_del=?", a.ArticleId, 0).Limit(1).Updates(value).Error; err != nil {
		return err
	}
	return nil
}

func (a ArticleTag) Delete(db *gorm.DB) error {

	if err := db.Where("id = ? AND is_del =?", a.Id, 0).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a ArticleTag) DeleteOne(db *gorm.DB) error {

	if err := db.Where("article_id = ? AND is_del =?", a.ArticleId, 0).Delete(&a).Limit(1).Error; err != nil {
		return err
	}
	return nil
}
