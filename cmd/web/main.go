package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/convox/docs/pkg/docs"
	"github.com/convox/docs/pkg/parser"
	"github.com/convox/docs/pkg/source"
	"github.com/convox/stdapi"
	"github.com/gobuffalo/packr"
)

var categorySlugs = []string{
	"introduction",
	"application",
	"development",
	"deployment",
	"management",
	"use-cases",
	"console",
	"enterprise",
	"example-apps",
	"reference",
	"external-services",
	"migration",
	"gen1",
	"help",
}

var documents docs.Documents

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	s := stdapi.New("docs", "docs.convox")

	// ds, err := docs.LoadGithub("convox", "convox", "docs", "master")
	// fmt.Printf("ds: %+v\n", ds)
	// fmt.Printf("err: %+v\n", err)

	fs, err := source.LoadS3("3.0.0.beta43")
	if err != nil {
		return err
	}

	ds, err := parser.Parse(fs)
	if err != nil {
		return err
	}

	documents = ds

	s.Router.Static("/assets/images", packr.NewBox("../../public/images"))
	s.Router.Static("/assets", packr.NewBox("../../public/assets"))

	s.Route("GET", "/", index)
	s.Route("GET", "/{slug:.*}", doc)

	// s.Route("GET", "/docs/{slug}", redirect)
	// s.Route("GET", "/docs/{slug}/", redirect)
	// s.Route("GET", "/{category}/{slug}", doc)
	// s.Route("GET", "/{category}/{slug}/", doc)
	// s.Route("GET", "/{slug}", redirect)
	// s.Route("GET", "/{slug}/", redirect)

	stdapi.LoadTemplates(packr.NewBox("../../templates"), helpers)

	// docs.Files = packr.NewBox("../../docs")

	// if err := docs.LoadCategories(categorySlugs...); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := docs.UploadIndex(); err != nil {
	// 	log.Printf("error: %s", err)
	// }

	if err := s.Listen("https", ":3000"); err != nil {
		return err
	}

	return nil
}

func helpers(c *stdapi.Context) template.FuncMap {
	return template.FuncMap{
		"env": func(s string) string {
			return os.Getenv(s)
		},
	}
}

func index(c *stdapi.Context) error {
	return c.Redirect(302, "/introduction/getting-started")
}

func doc(c *stdapi.Context) error {
	// cc, ok := docs.CategoryList().Find(c.Var("category"))
	// if !ok {
	// 	c.Response().WriteHeader(404)
	// 	params["Category"] = ""
	// 	return c.RenderTemplate("404", params)
	// }

	// params["Category"] = cc.Slug
	// params["CategoryName"] = cc.Name

	fmt.Printf("slug: %+v\n", c.Var("slug"))

	d, ok := documents.Find(c.Var("slug"))
	if !ok {
		c.Response().WriteHeader(404)
		return c.RenderTemplate("404", nil)
	}

	params := map[string]interface{}{
		"Document": template.HTML(d.Body),
		"Slug":     d.Slug,
	}

	// params["Document"] = template.HTML(d.Body)
	// params["Description"] = d.Description
	// params["Slug"] = d.Slug
	// params["Tags"] = d.Tags
	// params["Title"] = d.Title
	// params["URL"] = fmt.Sprintf("https://%s/%s/%s", c.Request().Host, cc.Slug, d.Slug)

	// if cc.Slug == "gen1" {
	// 	params["Deprecation"] = "Generation 1 has been deprecated and is not recommended for new applications."
	// }

	return c.RenderTemplate("doc", params)
}

// func redirect(c *stdapi.Context) error {
// 	for _, cc := range docs.CategoryList() {
// 		if d, ok := cc.Documents.Find(c.Var("slug")); ok {
// 			return c.Redirect(301, fmt.Sprintf("/%s/%s", cc.Slug, d.Slug))
// 		}
// 	}

// 	return stdapi.Errorf(404, "not found")
// }
