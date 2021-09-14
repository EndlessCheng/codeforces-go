package main

import (
	"html"
	"strings"
)

func entityParser(s string) string {
	return strings.ReplaceAll(html.UnescapeString(s), "⁄", "/")
}
