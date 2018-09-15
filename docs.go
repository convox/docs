package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/russross/blackfriday"
)

var (
	categories = Categories{}
	docs       = Documents{}
)

var categoryNames = map[string]string{
	"gen1": "Generation 1",
}

var (
	reDocument = regexp.MustCompile(`(?ms)(---([^-]*)---)?(.*)$`)
)

type Category struct {
	Documents Documents
	Name      string
	Slug      string
}

type Categories []Category

type Document struct {
	Body  []byte
	Order int
	Slug  string
	Title string
}

type Documents []Document

func LoadCategories(slugs ...string) error {
	for _, slug := range slugs {
		if err := LoadCategory(slug); err != nil {
			return err
		}
	}

	return nil
}

func LoadCategory(slug string) error {
	root := filepath.Join("docs", slug)

	tokens := strings.Split(slug, "-")

	for i, token := range tokens {
		tokens[i] = strings.Title(token)
	}

	title := strings.Join(tokens, " ")

	if t := categoryNames[slug]; t != "" {
		title = t
	}

	c := Category{
		Name:      title,
		Slug:      slug,
		Documents: Documents{},
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		if info == nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		// category := filepath.Dir(rel)
		name := filepath.Base(rel)
		name = strings.TrimSuffix(name, ".md")

		slug = strings.Replace(name, ".", "-", -1)

		d := Document{
			Slug: slug,
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		m := reDocument.FindSubmatch(data)

		if len(m) != 4 {
			return nil
		}

		var front map[string]string

		if err := yaml.Unmarshal(m[1], &front); err != nil {
			return err
		}

		d.Title = front["title"]

		if d.Title == "" {
			d.Title = name
		}

		d.Order = 50000

		if os, ok := front["order"]; ok {
			o, err := strconv.Atoi(os)
			if err != nil {
				return err
			}
			d.Order = o
		}

		parsed := blackfriday.Run(m[3],
			blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.AutoHeadingIDs),
			blackfriday.WithRenderer(blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{Flags: blackfriday.CommonHTMLFlags})),
		)
		// blackfriday.WithExtensions(blackfriday.LaxHTMLBlocks),

		d.Body = parsed

		c.Documents = append(c.Documents, d)

		docs = append(docs, d)

		return nil
	})
	if err != nil {
		return err
	}

	sort.Slice(c.Documents, c.Documents.Less)

	categories = append(categories, c)

	return nil
}

func (ds Documents) Find(slug string) (*Document, bool) {
	for _, d := range ds {
		if d.Slug == slug {
			return &d, true
		}
	}

	return nil, false
}

func (ds Documents) Less(i, j int) bool {
	if ds[i].Order == ds[j].Order {
		return ds[i].Title < ds[j].Title
	}
	return ds[i].Order < ds[j].Order
}

// func loadDocs() error {
//   root := filepath.Join(wd, "docs")

//   fmt.Printf("root = %+v\n", root)

// }
