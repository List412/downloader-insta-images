package files

import (
	"downloader-insta-images/internal/network"
	"io"
	"os"
	"path/filepath"
)

func Save(url string, path string) error {
	err := createPath(path)
	if err != nil {
		return err
	}
	response, err := network.Get(url)
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func createPath(path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		return err
	}
	return nil
}
