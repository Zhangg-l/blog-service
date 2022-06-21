package dao

import (
	"go_code/project8/blog-service/internal/model"
	"go_code/project8/blog-service/pkg/app"
)

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}

	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}

	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{Name: name, State: state, Model: &model.Model{CreatedBy: createdBy}}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(name string, state uint8, id uint32, modifiedBy string) error {
	tag := model.Tag{Name: name, State: state, Model: &model.Model{Id: id, ModifiedBy: modifiedBy}}
	values := map[string]interface{}{
		"state":       0,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{Id: id}}
	return tag.Delete(d.engine)
}

func (d *Dao) GetTag(id uint32, state uint8) (*model.Tag, error) {
	var tag = model.Tag{Model: &model.Model{Id: id}, State: state}
	return tag.GetTag(d.engine)
}
