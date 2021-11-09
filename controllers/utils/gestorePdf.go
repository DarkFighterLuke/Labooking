package utils

import (
	"Labooking/models/utils"
	"github.com/beego/beego/v2/server/web"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
)

func SaveUploadedPdf(file multipart.File, pathAppConfig string) (fileName string, err error) {
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	path, err := web.AppConfig.String(pathAppConfig)
	if err != nil {
		return "", err
	}

	fileName = utils.RandStringRunes(32)
	filePath := path + fileName + ".pdf"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			return "", err
		}
	}

	err = ioutil.WriteFile(filePath, fileBytes, 0655)
	if err != nil {
		return "", err
	}
	return fileName, err
}
