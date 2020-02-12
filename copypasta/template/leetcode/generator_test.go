package leetcode

import (
	"os"
	"testing"
)

func TestGenLeetCodeTests(t *testing.T) {
	username := os.Getenv("LEETCODE_USERNAME")
	password := os.Getenv("LEETCODE_PASSWORD")
	if err := GenLeetCodeTests(username, password); err != nil {
		t.Fatal(err)
	}
}
