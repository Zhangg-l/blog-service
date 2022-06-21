package model

import "github.com/jinzhu/gorm"

type Article struct {
	*Model
	Title         string `json:"title"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	Desc          string `json:"desc"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_Article"
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Updates(values).Where("id=? AND is_del=?", a.Id, 0).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db.Where("id=? AND state = ? AND is_del=?", a.Id, a.State, 0)
	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	return article, nil
}

func (a Article) Delete(db *gorm.DB) error {
	if err := db.Where("id=? AND is_del=?", a.Id, 0).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

// 获取关联信息
type ArticleRow struct {
	ArticleId     uint32
	TagId         uint32
	TagName       string
	ArticleTitle  string
	ArticleDesc   string
	CoverImageUrl string
	Content       string
}

func (a Article) ListByTagId(db *gorm.DB, tagId uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{"ar.id AS article_id", "ar.title AS article_title", "ar.desc AS article_desc", "ar.cover_image_url",
		"ar.content"}
	fields = append(fields, []string{"t.id AS tag_id", "t.name AS tag_name"}...)

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(fields).Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN`"+Tag{}.TableName()+"`AS t ON at.tag_id=t.id").
		Joins("LEFT JOIN`"+Article{}.TableName()+"`AS ar ON at.article_id=ar.id").
		Where("at.`tag_id`=? AND ar.state=? AND ar.is_del=?", tagId, a.State, 0).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var articleRow []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		if err := rows.Scan(&r.ArticleId, &r.ArticleTitle, &r.ArticleDesc,
			&r.CoverImageUrl, &r.Content, &r.TagId, &r.TagName); err != nil {
			return nil, err
		}
		articleRow = append(articleRow, r)
	}
	return articleRow, nil
}
func (a Article) CountByTagId(db *gorm.DB, tagId uint32) (int, error) {
	var Count int

	err := db.Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id=t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id=ar.id").
		Where("at.`tag_id`=? AND ar.state=? AND ar.is_del=?", tagId, a.State, 0).Count(&Count).Error
	if err != nil {
		return 0, err
	}
	return Count, nil

}
