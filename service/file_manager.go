package service

import (
	"errors"
	"github.com/fghwett/typecho-to-halo/config"
	"github.com/imroc/req/v3"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type FileManager struct {
	conf *config.FileManager
}

func NewFileManager(conf *config.FileManager) (*FileManager, error) {

	fm := &FileManager{
		conf: conf,
	}

	_, err := os.Stat(conf.Dir)

	if err == nil {
		return fm, nil
	}

	if !os.IsNotExist(err) {
		return nil, err
	}

	err = os.MkdirAll(conf.Dir, os.ModePerm)

	return fm, err
}

func (fileManager *FileManager) DownFile(link string) (*string, error) {
	name := link[strings.LastIndex(link, "/")+1:]

	dir := filepath.Join(fileManager.conf.Dir, name)

	if _, err := os.Stat(dir); err == nil {
		return &dir, nil
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	resp, err := req.C().R().SetOutputFile(dir).Get(link)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("文件下载异常")
	}

	return &dir, nil
}

func (fileManager *FileManager) DeleteFile(filename string) error {
	return os.Remove(filename)
}
