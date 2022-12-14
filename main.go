package main

import (
	"encoding/xml"
	"io"
	"os"
	"text/template"

	"github.com/ksrnnb/saml/route"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

const httpPostBinding = "urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"

func main() {
	f, err := os.ReadFile("./metadata.xml")
	if err != nil {
		panic(err)
	}

	md := Metadata{}
	xml.Unmarshal(f, &md.EntityDescriptor)

	t := &Template{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	route.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":3000"))

}
