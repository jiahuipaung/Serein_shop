package upload

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	conf "serein/config"
	util "serein/pkg/utils/log"
)

//
func ProductUploadToLocalStatic(file multipart.File, productName string) (filepath string, err error) {
	basePath := "." + conf.Config.PhotoPath.ProductPath + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := fmt.Sprintf("%s%s.jpg", basePath, productName)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		util.LogrusObj.Error(err)
		return "", err
	}
	err = ioutil.WriteFile(productPath, content, 0666)
	if err != nil {
		util.LogrusObj.Error(err)
		return "", err
	}

	return fmt.Sprint("product/%s.jpg", productName), err
}

// DirExistOrNot 判断文件是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

// CreateDir 创建文件夹
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}