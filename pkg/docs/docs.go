package docs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/go-yaml/yaml"
	"github.com/russross/blackfriday"
)

var (
	algolia    algoliasearch.Index
	categories = Categories{}
	documents  = Documents{}
)

var categoryNames = map[string]string{
	"application": "Application Setup",
	"gen1":        "Generation 1",
	"migration":   "Migration Guides",
}

var (
	reDocument    = regexp.MustCompile(`(?ms)(---(.*?)---)?(.*)$`)
	reMarkdownDiv = regexp.MustCompile(`(?ms)(<div.*?markdown="1".*?>(.*?)</div>)`)
	reTag         = regexp.MustCompile(`(?ms)(<[\w]+.*?>(.*?)</[\w]+>)`)
)

type Category struct {
	Documents Documents
	Name      string
	Slug      string
}

type Categories []Category

type Document struct {
	Body        []byte
	Description string
	Order       int
	Slug        string
	Tags        []string
	Title       string
}

type Documents []Document

func init() {
	algolia = algoliasearch.NewClient(os.Getenv("ALGOLIA_APP"), os.Getenv("ALGOLIA_KEY_ADMIN")).InitIndex(os.Getenv("ALGOLIA_INDEX"))
}

func CategoryList() Categories {
	return categories
}

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
		if err != nil {
			return err
		}

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

		d.Description = front["description"]
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

		d.Tags = []string{}

		for _, t := range strings.Split(front["tags"], ",") {
			if tt := strings.TrimSpace(t); tt != "" {
				d.Tags = append(d.Tags, tt)
			}
		}

		sort.Strings(d.Tags)

		markdown := m[3]

		for _, n := range reMarkdownDiv.FindAllSubmatch(markdown, -1) {
			np := blackfriday.Run(n[2],
				blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.AutoHeadingIDs|blackfriday.LaxHTMLBlocks),
			)

			markdown = bytes.Replace(markdown, n[2], np, -1)
		}

		parsed := blackfriday.Run(markdown,
			blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.AutoHeadingIDs|blackfriday.LaxHTMLBlocks),
		)

		d.Body = parsed

		c.Documents = append(c.Documents, d)

		documents = append(documents, d)

		return nil
	})
	if err != nil {
		return err
	}

	sort.Slice(c.Documents, c.Documents.Less)

	categories = append(categories, c)

	return nil
}

func UploadIndex() error {
	algolia.Clear()

	os := []algoliasearch.Object{}

	for _, c := range categories {
		for _, d := range c.Documents {
			body := d.Body

			for {
				stripped := reTag.ReplaceAll(body, []byte("$2"))
				if bytes.Equal(body, stripped) {
					break
				}
				body = stripped
			}

			os = append(os, algoliasearch.Object{
				"objectID":       fmt.Sprintf("%s:%s", c.Slug, d.Slug),
				"category_slug":  c.Slug,
				"category_title": c.Name,
				"body":           string(body),
				"slug":           string(d.Slug),
				"title":          string(d.Title),
				"url":            fmt.Sprintf("/%s/%s", c.Slug, d.Slug),
			})
		}
	}

	if _, err := algolia.AddObjects(os); err != nil {
		return err
	}

	return nil
}

func (cs Categories) Find(slug string) (*Category, bool) {
	for _, c := range cs {
		if c.Slug == slug {
			return &c, true
		}
	}

	return nil, false
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
