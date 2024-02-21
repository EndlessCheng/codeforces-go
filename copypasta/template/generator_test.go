package main

import (
	"os"
	"strings"
	"testing"
)

func TestGenCodeforcesProblemTemplates(t *testing.T) {
	raw, err := os.ReadFile("data-codeforces.txt")
	if err != nil {
		t.Fatal(err)
	}
	problemURL := strings.TrimSpace(string(raw))
	if err := GenCodeforcesProblemTemplates(problemURL, true); err != nil {
		t.Fatal(err)
	}
}
