package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/convox/stdapi"
)

var categorySlugs = []string{
	"introduction",
	"deployment",
	"development",
	"management",
	"monitoring",
	"resources",
	"integrations",
	"console",
	"reference",
	"gen1",
}

func init() {
	if err := LoadCategories(categorySlugs...); err != nil {
		log.Fatal(err)
	}
}

func main() {
	s := stdapi.New("docs", "docs.convox")

	s.Router.Static("/assets/", "./assets")

	s.Route("GET", "/", index)
	s.Route("GET", "/{category}/{slug}", doc)
	s.Route("GET", "/{slug}", redirect)

	stdapi.LoadTemplates("./templates", helpers)

	if err := s.Listen("https", ":3000"); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}

func helpers(c *stdapi.Context) template.FuncMap {
	return template.FuncMap{}
}

func index(c *stdapi.Context) error {
	return c.RenderTemplate("doc", nil)
}

func doc(c *stdapi.Context) error {
	d, ok := docs.Find(c.Var("slug"))
	if !ok {
		return stdapi.Errorf(404, "not found")
	}

	params := map[string]interface{}{
		"Categories": categories,
		"Document":   template.HTML(d.Body),
		"Slug":       c.Var("slug"),
		"Title":      d.Title,
	}

	return c.RenderTemplate("doc", params)
}

func redirect(c *stdapi.Context) error {
	return nil
}
