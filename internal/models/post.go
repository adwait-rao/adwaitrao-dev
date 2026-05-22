package models

import "html/template"

// Metadata parsed from the top of the Markdown file
type FrontMatter struct {
	Title       string `yaml:"title"`
	Date        string `yaml:"date"`
	Description string `yaml:"description"`
}

type Post struct {
	FrontMatter
	Slug    string
	Content template.HTML
}
