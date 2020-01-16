package source

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func LoadTarball(url, prefix string, skip int) (Files, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	gz, err := gzip.NewReader(res.Body)
	if err != nil {
		return nil, err
	}

	t := tar.NewReader(gz)

	fs := Files{}

	for {
		h, err := t.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		path := strings.Join(strings.Split(h.Name, "/")[skip:], "/")

		if strings.HasPrefix(path, prefix) {
			data, err := ioutil.ReadAll(t)
			if err != nil {
				return nil, err
			}

			fs = append(fs, File{
				Path: strings.TrimPrefix(path, prefix),
				Body: data,
			})
		}
	}

	return fs, nil
}
