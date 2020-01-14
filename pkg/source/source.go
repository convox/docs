package source

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type File struct {
	Body []byte
	Path string
}

type Files []File

func LoadS3(version string) (Files, error) {
	fs := Files{}

	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	s3c := s3.New(s)

	bucket := os.Getenv("BUCKET")
	prefix := fmt.Sprintf("%s/", version)

	req := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	}

	fmt.Printf("req: %+v\n", req)

	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			r, _ := s3c.ListObjectsV2Request(req)
			r.SetContext(context.Background())
			return r, nil
		},
	}

	for p.Next() {
		for _, obj := range p.Page().(*s3.ListObjectsV2Output).Contents {
			res, err := s3c.GetObject(&s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    obj.Key,
			})
			if err != nil {
				return nil, err
			}

			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}
			res.Body.Close()

			fs = append(fs, File{
				Body: data,
				Path: strings.TrimPrefix(*obj.Key, prefix),
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return fs, nil
}
