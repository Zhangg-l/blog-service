package model

import (
	"fmt"
	"go_code/project8/blog-service/global"
	"go_code/project8/blog-service/pkg/setting"
	"time"

	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	Id         uint32 `gorm:"primary_key" json:"id"`
	CreatedOn  uint32 `  json:"created_on"`
	DeletedOn  uint32 `  json:"deleted_on"`
	ModifiedOn uint32 `  json:"modified_on"`
	ModifiedBy string `  json:"modified_by"`
	CreatedBy  string `  json:"created_by"`
}

func NewDBEngine(DatabaseSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	db, err := gorm.Open(DatabaseSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			DatabaseSetting.Username,
			DatabaseSetting.Password,
			DatabaseSetting.Host,
			DatabaseSetting.DBName,
			DatabaseSetting.Charset,
			DatabaseSetting.ParseTime,
		))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	// 注册回调行为
	db.Callback().Create().Replace("gorm:update_time_stamp",
		updateTimeStampCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp",
		updateTimeStampUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	db.DB().SetMaxIdleConns(DatabaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(DatabaseSetting.MaxOpenConns)
	otgorm.AddGormCallbacks(db)
	return db, nil
}

// 编写回调代码
func updateTimeStampCreateCallback(scope *gorm.Scope) {

	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createdTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createdTimeField.IsBlank {
				_ = createdTimeField.Set(nowTime)
			}
		}

		if modifiedTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifiedTimeField.IsBlank {
				_ = modifiedTimeField.Set(nowTime)
			}
		}

	}
}
func updateTimeStampUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_cloumn"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		isDelField, hasIsDelField := scope.FieldByName("IsDel")

		if scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(
				fmt.Sprintf("DELETE FROM %v%v%v",
					scope.QuotedTableName(),
					addExtraSpaceIfExist(scope.CombinedConditionSql()),
					addExtraSpaceIfExist(extraOption),
				)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
