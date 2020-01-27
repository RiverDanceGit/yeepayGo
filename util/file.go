package util

import (
	"io/ioutil"
	"os"
)

func FileGetContents(fileName string) (string, error) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func FilePutContents(fileName string, content string) (int, error) {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		return 0, err
	}
	return file.Write([]byte(content))
}
