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
	if err := GenCodeforcesProblemTemplates(problemURL, false); err != nil {
		t.Fatal(err)
	}
}
