package source

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func LoadLocal(path string) (Files, error) {
	fs := Files{}

	err := filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		fs = append(fs, File{
			Path: strings.TrimPrefix(file, path),
			Body: data,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return fs, nil
}
