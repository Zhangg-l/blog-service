package upload

import (
	"go_code/project8/blog-service/global"
	"go_code/project8/blog-service/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

// 得到文件后缀
func GetFileExts(fileName string) string {
	// 返回后缀
	ext := path.Ext(fileName)
	return ext
}

// 得到文件加密后的名字
func GetFileName(name string) string {

	ext := GetFileExts(name)
	fileName := strings.TrimSuffix(name, ext)
	return util.EncodingMD5(fileName) + ext
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func CheckSavePath(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

// 检查权限
func CheckPremession(path string) bool {
	_, err := os.Stat(path)
	return os.IsPermission(err)
}

// 检查文件后缀是否包含约定的后缀
func CheckContainExt(fType FileType, name string) bool {
	ext := GetFileExts(name)
	switch fType {
	case TypeImage:
		for _, val := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(ext) == strings.ToUpper(val) {
				return true
			}
		}
	}
	return false
}

func CheckMaxSize(fType FileType, f multipart.File) bool {
	bt, err := ioutil.ReadAll(f)
	if err != nil {
		return false
	}
	bLen := len(bt)
	switch fType {
	case TypeImage:
		if bLen >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	return err
}

// 把file的数据copy 到dst中
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err
}
