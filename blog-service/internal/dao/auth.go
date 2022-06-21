package dao

import "go_code/project8/blog-service/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}

	return auth.Get(d.engine)
}
