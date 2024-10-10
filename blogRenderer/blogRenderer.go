package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

var (
	postTemplate embed.FS // embedding a file system
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(w io.Writer, post Post) error {
	// templ, err := template.New("blog").Parse(postTemplate) // Parse parses the template string, identifies the placeholders

	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, post); err != nil { // Execute will substitute placeholders with data passed in (Post)
		return err
	}

	return nil
}
