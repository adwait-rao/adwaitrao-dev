package parser

import (
	"bytes"
	"errors"
	"strings"

	"github.com/adwait-rao/adwaitrao-dev/internal/models"
	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
)

func ParseMarkdownFile(rawContent []byte) (models.FrontMatter, string, error) {
	var fm models.FrontMatter
	contentStr := string(rawContent)

	if !strings.HasPrefix(contentStr, "---\n") {
		return fm, "", errors.New("missing front matter at the start of file")
	}

	parts := strings.SplitN(contentStr, "---\n", 3)
	if len(parts) < 3 {
		return fm, "", errors.New("malformed front matter: missing closing delimiter")
	}

	yamlBlock := parts[1]
	markdownBody := parts[2]

	if err := yaml.Unmarshal([]byte(yamlBlock), &fm); err != nil {
		return fm, "", err
	}

	return fm, markdownBody, nil
}

func ConvertMarkdownToHTML(markdown string) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}
