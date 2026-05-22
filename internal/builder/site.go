package builder

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/adwait-rao/adwaitrao-dev/internal/models"
	"github.com/adwait-rao/adwaitrao-dev/internal/parser"
)

// BuildSite reads from srcDir, processes Markdown files, and outputs static HTML to outDir.
func BuildSite(srcDir, outDir, templateDir string) error {
	if err := os.RemoveAll(outDir); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(outDir, "blog"), 0755); err != nil {
		return err
	}

	baseTmpl, err := template.ParseFiles(
		filepath.Join(templateDir, "base.html"),
		filepath.Join(templateDir, "blog.html"),
	)
	if err != nil {
		return err
	}

	indexTmpl, err := template.ParseFiles(
		filepath.Join(templateDir, "base.html"),
		filepath.Join(templateDir, "index.html"),
	)
	if err != nil {
		return err
	}

	files, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	var allPosts []models.Post

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".md" {
			continue // Skip folders or non-markdown files
		}

		rawPath := filepath.Join(srcDir, file.Name())
		rawBytes, err := os.ReadFile(rawPath)
		if err != nil {
			return err
		}

		fm, mdBody, err := parser.ParseMarkdownFile(rawBytes)
		if err != nil {
			return err
		}

		htmlBody, err := parser.ConvertMarkdownToHTML(mdBody)
		if err != nil {
			return err
		}

		slug := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		post := models.Post{
			FrontMatter: fm,
			Slug:        slug,
			Content:     template.HTML(htmlBody),
		}

		allPosts = append(allPosts, post)

		outPath := filepath.Join(outDir, "blog", slug+".html")
		outFile, err := os.Create(outPath)
		if err != nil {
			return err
		}

		if err := baseTmpl.ExecuteTemplate(outFile, "base", post); err != nil {
			outFile.Close()
			return err
		}
		outFile.Close()
	}

	indexFile, err := os.Create(filepath.Join(outDir, "index.html"))
	if err != nil {
		return err
	}
	defer indexFile.Close()

	return indexTmpl.ExecuteTemplate(indexFile, "base", allPosts)
}
