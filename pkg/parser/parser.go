package parser

import (
	"bytes"
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/convox/docs/pkg/docs"
	"github.com/convox/docs/pkg/source"
	"github.com/russross/blackfriday"
	yaml "gopkg.in/yaml.v2"
)

// var (
// 	categories = Categories{}
// 	documents  = Documents{}
// )

var categoryNames = map[string]string{
	"application": "Application Setup",
	"gen1":        "Generation 1",
	"migration":   "Migration Guides",
}

var (
	reDocument    = regexp.MustCompile(`(?ms)(---(.*?)---)?(.*)$`)
	reTitle       = regexp.MustCompile(`<h1.*?>(.*?)</h1>`)
	reLink        = regexp.MustCompile(`href="([^"]+)"`)
	reMarkdownDiv = regexp.MustCompile(`(?ms)(<div.*?markdown="1".*?>(.*?)</div>)`)
)

// var (
// 	Files packr.Box
// )

// type Category struct {
// 	Documents Documents
// 	Name      string
// 	Slug      string
// }

// type Categories []Category

// type Document struct {
// 	Body        []byte
// 	Description string
// 	Order       int
// 	Slug        string
// 	Tags        []string
// 	Title       string
// }

// type Documents []Document

// type Version struct {
// 	Categories Categories
// 	Documents  Documents
// 	Name       string
// }

func Parse(files source.Files) (*docs.Documents, error) {
	ds := docs.NewDocuments()

	for _, f := range files {
		d, err := parseDocument(f.Path, f.Body)
		if err != nil {
			return nil, err
		}

		if err := ds.Add(*d); err != nil {
			return nil, err
		}
	}

	return &ds, nil
}

func parseDocument(path string, data []byte) (*docs.Document, error) {
	name := strings.TrimSuffix(path, ".md")
	readme := strings.HasSuffix(path, "/README.md")
	slug := strings.TrimSuffix(strings.Replace(name, ".", "-", -1), "/README")

	d := &docs.Document{
		Slug: slug,
	}

	m := reDocument.FindSubmatch(data)

	if len(m) != 4 {
		return nil, nil
	}

	var front map[string]string

	if err := yaml.Unmarshal(m[1], &front); err != nil {
		return nil, err
	}

	d.Description = front["description"]
	d.Title = front["title"]

	d.Order = 50000

	if os, ok := front["order"]; ok {
		o, err := strconv.Atoi(os)
		if err != nil {
			return nil, err
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

	d.Body = reLink.ReplaceAllFunc(parsed, func(link []byte) []byte {
		if match := reLink.FindSubmatch(link); len(match) < 2 {
			return link
		} else {
			u, err := url.Parse(string(match[1]))
			if err != nil {
				return link
			}
			if u.Host == "" {
				u.Path = strings.TrimSuffix(u.Path, ".md")
				if readme {
					u.Path = fmt.Sprintf("/%s/%s", d.Slug, u.Path)
				}
			}
			return []byte(fmt.Sprintf("href=%q", u.String()))
		}
	})

	if d.Title == "" {
		if m := reTitle.FindStringSubmatch(string(d.Body)); len(m) > 1 {
			d.Title = m[1]
		}
	}

	if d.Title == "" {
		parts := strings.Split(d.Slug, "/")
		d.Title = strings.Title(strings.ReplaceAll(parts[len(parts)-1], "-", " "))
	}

	return d, nil
}

// func CategoryList() Categories {
// 	return categories
// }

// func LoadCategories(slugs ...string) error {
// 	for _, slug := range slugs {
// 		if err := LoadCategory(slug); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func LoadCategory(slug string) error {
// 	// root := filepath.Join("docs", slug)

// 	tokens := strings.Split(slug, "-")

// 	for i, token := range tokens {
// 		tokens[i] = strings.Title(token)
// 	}

// 	title := strings.Join(tokens, " ")

// 	if t := categoryNames[slug]; t != "" {
// 		title = t
// 	}

// 	c := Category{
// 		Name:      title,
// 		Slug:      slug,
// 		Documents: Documents{},
// 	}

// 	err := Files.WalkPrefix(fmt.Sprintf("%s/", slug), func(path string, file packd.File) error {
// 		name := filepath.Base(path)
// 		name = strings.TrimSuffix(name, ".md")

// 		slug = strings.Replace(name, ".", "-", -1)

// 		d := Document{
// 			Slug: slug,
// 		}

// 		data, err := ioutil.ReadAll(file)
// 		if err != nil {
// 			return err
// 		}

// 		m := reDocument.FindSubmatch(data)

// 		if len(m) != 4 {
// 			return nil
// 		}

// 		var front map[string]string

// 		if err := yaml.Unmarshal(m[1], &front); err != nil {
// 			return err
// 		}

// 		d.Description = front["description"]
// 		d.Title = front["title"]

// 		if d.Title == "" {
// 			d.Title = name
// 		}

// 		d.Order = 50000

// 		if os, ok := front["order"]; ok {
// 			o, err := strconv.Atoi(os)
// 			if err != nil {
// 				return err
// 			}
// 			d.Order = o
// 		}

// 		d.Tags = []string{}

// 		for _, t := range strings.Split(front["tags"], ",") {
// 			if tt := strings.TrimSpace(t); tt != "" {
// 				d.Tags = append(d.Tags, tt)
// 			}
// 		}

// 		sort.Strings(d.Tags)

// 		markdown := m[3]

// 		for _, n := range reMarkdownDiv.FindAllSubmatch(markdown, -1) {
// 			np := blackfriday.Run(n[2],
// 				blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.AutoHeadingIDs|blackfriday.LaxHTMLBlocks),
// 			)

// 			markdown = bytes.Replace(markdown, n[2], np, -1)
// 		}

// 		parsed := blackfriday.Run(markdown,
// 			blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.AutoHeadingIDs|blackfriday.LaxHTMLBlocks),
// 		)

// 		d.Body = parsed

// 		c.Documents = append(c.Documents, d)

// 		documents = append(documents, d)

// 		return nil
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	sort.Slice(c.Documents, c.Documents.Less)

// 	categories = append(categories, c)

// 	return nil
// }

// func (cs Categories) Find(slug string) (*Category, bool) {
// 	for _, c := range cs {
// 		if c.Slug == slug {
// 			return &c, true
// 		}
// 	}

// 	return nil, false
// }

// func (ds Documents) Find(slug string) (*Document, bool) {
// 	for _, d := range ds {
// 		if d.Slug == slug {
// 			return &d, true
// 		}
// 	}

// 	return nil, false
// }

// func (ds Documents) Less(i, j int) bool {
// 	if ds[i].Order == ds[j].Order {
// 		return ds[i].Title < ds[j].Title
// 	}
// 	return ds[i].Order < ds[j].Order
// }
