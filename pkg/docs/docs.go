package docs

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

var (
	categories = Categories{
		Category{"getting-started", "Getting Started"},
		Category{"tutorials", "Tutorials"},
		Category{"installation", "Installation"},
		Category{"configuration", "Configuration"},
		Category{"development", "Development"},
		Category{"deployment", "Deployment"},
		Category{"management", "Management"},
		Category{"reference", "Reference"},
		Category{"integrations", "Integrations"},
		Category{"help", "Help"},
	}
)

var (
	reTag = regexp.MustCompile(`(?ms)(<[\w]+.*?>(.*?)</[\w]+>)`)
)

type Category struct {
	Slug string
	Name string
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

type Documents struct {
	documents []Document
}

func NewDocuments() Documents {
	return Documents{
		documents: []Document{},
	}
}

func (cs Categories) Find(slug string) Category {
	for _, c := range cs {
		if c.Slug == slug {
			return c
		}
	}
	return Category{}
}

func (ds *Documents) Add(d Document) error {
	ds.documents = append(ds.documents, d)
	return nil
}

func (ds *Documents) Categories() Categories {
	return categories
}

func (ds *Documents) Children(slug string) []Document {
	docs := []Document{}

	for _, d := range ds.documents {
		if strings.HasPrefix(d.Slug, fmt.Sprintf("%s/", slug)) && level(d.Slug) == level(slug)+1 {
			docs = append(docs, d)
		}
	}

	sort.Slice(docs, func(i, j int) bool {
		if docs[i].Order == docs[j].Order {
			return docs[i].Title < docs[j].Title
		}
		return docs[i].Order < docs[j].Order
	})

	return docs
}

func (ds *Documents) Find(slug string) (*Document, bool) {
	for _, d := range ds.documents {
		if d.Slug == slug {
			return &d, true
		}
	}

	return nil, false
}

func (ds *Documents) UploadIndex() error {
	if os.Getenv("ALGOLIA_APP") == "" {
		return nil
	}

	algolia := algoliasearch.NewClient(os.Getenv("ALGOLIA_APP"), os.Getenv("ALGOLIA_KEY_ADMIN")).InitIndex(os.Getenv("ALGOLIA_INDEX"))

	algolia.Clear()

	os := []algoliasearch.Object{}

	for _, d := range ds.documents {
		body := d.Body

		for {
			stripped := reTag.ReplaceAll(body, []byte("$2"))
			if bytes.Equal(body, stripped) {
				break
			}
			body = stripped
		}

		if len(body) > 8000 {
			body = body[0:8000]
		}

		c := categories.Find(strings.Split(d.Slug, "/")[0])

		os = append(os, algoliasearch.Object{
			"objectID":       d.Slug,
			"category_slug":  c.Slug,
			"category_title": strings.Join(ds.Breadcrumbs(d.Slug), " » "),
			"body":           string(body),
			"slug":           string(d.Slug),
			"title":          string(d.Title),
			"url":            "/" + d.Slug,
		})
	}

	if _, err := algolia.AddObjects(os); err != nil {
		return err
	}

	return nil
}

func (ds *Documents) Breadcrumbs(slug string) []string {
	d, ok := ds.Find(slug)
	if !ok {
		return []string{}
	}

	bs := []string{}

	bs = append(bs, d.Category().Name)

	parts := strings.Split(d.Slug, "/")

	for i := 1; i < len(parts); i++ {
		if d, ok := ds.Find(strings.Join(parts[0:i], "/")); ok {
			bs = append(bs, d.Title)
		}
	}

	return bs
}

func (d *Document) Category() Category {
	return categories.Find(strings.Split(d.Slug, "/")[0])
}

func (d *Document) Level() int {
	return level(d.Slug)
}

func level(slug string) int {
	return len(strings.Split(slug, "/")) - 1
}

// func (ds Documents) Less(i, j int) bool {
// 	if ds.documents[i].Order == ds.docuements[j].Order {
// 		return ds[i].Title < ds[j].Title
// 	}
// 	return ds[i].Order < ds[j].Order
// }

// func (ds *Documents) Hierarchy(category, slug string) []Hierarchy {
// 	hs := []Hierarchy{}

// 	fmt.Printf("category: %+v\n", category)
// 	fmt.Printf("slug: %+v\n", slug)

// 	for _, d := range ds.documents {
// 		if !strings.HasPrefix(d.Slug, fmt.Sprintf("%s/", category)) {
// 			continue
// 		}

// 		if len(strings.Split(d.Slug, "/")) < 2 && !strings.HasPrefix(d.Slug, slug) {
// 			continue
// 		}

// 		hs = append(hs, Hierarchy{
// 			Document: d,
// 			Children: ds.Hierarchy(category, d.Slug),
// 		})
// 	}

// 	return hs
// }

// var (
// 	categories = Categories{}
// 	documents  = Documents{}
// )

// var categoryNames = map[string]string{
// 	"application": "Application Setup",
// 	"gen1":        "Generation 1",
// 	"migration":   "Migration Guides",
// }

// var (
// 	reDocument    = regexp.MustCompile(`(?ms)(---(.*?)---)?(.*)$`)
// 	reMarkdownDiv = regexp.MustCompile(`(?ms)(<div.*?markdown="1".*?>(.*?)</div>)`)
// 	reTag         = regexp.MustCompile(`(?ms)(<[\w]+.*?>(.*?)</[\w]+>)`)
// )

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

// func Parse(files source.Files) (Documents, error) {
// 	ds := Documents{}

// 	for _, f := range files {
// 		d, err := parseDocument(f.Path, f.Body)
// 		if err != nil {
// 			return nil, err
// 		}

// 		ds = append(ds, *d)
// 	}

// 	return ds, nil
// }

// func parseDocument(path string, data []byte) (*Document, error) {
// 	name := strings.TrimSuffix(filepath.Base(path), ".md")
// 	slug := strings.Replace(name, ".", "-", -1)

// 	d := &Document{
// 		Slug: slug,
// 	}

// 	m := reDocument.FindSubmatch(data)

// 	if len(m) != 4 {
// 		return nil, nil
// 	}

// 	var front map[string]string

// 	if err := yaml.Unmarshal(m[1], &front); err != nil {
// 		return nil, err
// 	}

// 	d.Description = front["description"]
// 	d.Title = front["title"]

// 	if d.Title == "" {
// 		d.Title = name
// 	}

// 	d.Order = 50000

// 	if os, ok := front["order"]; ok {
// 		o, err := strconv.Atoi(os)
// 		if err != nil {
// 			return nil, err
// 		}
// 		d.Order = o
// 	}

// 	d.Tags = []string{}

// 	for _, t := range strings.Split(front["tags"], ",") {
// 		if tt := strings.TrimSpace(t); tt != "" {
// 			d.Tags = append(d.Tags, tt)
// 		}
// 	}

// 	sort.Strings(d.Tags)

// 	markdown := m[3]

// 	for _, n := range reMarkdownDiv.FindAllSubmatch(markdown, -1) {
// 		np := blackfriday.Run(n[2],
// 			blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.AutoHeadingIDs|blackfriday.LaxHTMLBlocks),
// 		)

// 		markdown = bytes.Replace(markdown, n[2], np, -1)
// 	}

// 	parsed := blackfriday.Run(markdown,
// 		blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.AutoHeadingIDs|blackfriday.LaxHTMLBlocks),
// 	)

// 	d.Body = parsed

// 	return d, nil
// }

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

// func UploadIndex() error {
// 	if os.Getenv("ALGOLIA_APP") == "" {
// 		return nil
// 	}

// 	algolia := algoliasearch.NewClient(os.Getenv("ALGOLIA_APP"), os.Getenv("ALGOLIA_KEY_ADMIN")).InitIndex(os.Getenv("ALGOLIA_INDEX"))

// 	algolia.Clear()

// 	os := []algoliasearch.Object{}

// 	for _, c := range categories {
// 		for _, d := range c.Documents {
// 			body := d.Body

// 			for {
// 				stripped := reTag.ReplaceAll(body, []byte("$2"))
// 				if bytes.Equal(body, stripped) {
// 					break
// 				}
// 				body = stripped
// 			}

// 			if len(body) > 8000 {
// 				body = body[0:8000]
// 			}

// 			os = append(os, algoliasearch.Object{
// 				"objectID":       fmt.Sprintf("%s:%s", c.Slug, d.Slug),
// 				"category_slug":  c.Slug,
// 				"category_title": c.Name,
// 				"body":           string(body),
// 				"slug":           string(d.Slug),
// 				"title":          string(d.Title),
// 				"url":            fmt.Sprintf("/%s/%s", c.Slug, d.Slug),
// 			})
// 		}
// 	}

// 	if _, err := algolia.AddObjects(os); err != nil {
// 		return err
// 	}

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
