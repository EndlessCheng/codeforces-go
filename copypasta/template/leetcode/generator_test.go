package leetcode

import (
	"os"
	"testing"
)

// TODO: 确认是否登录以及默认语言是否正确
func TestGenLeetCodeTests(t *testing.T) {
	username := os.Getenv("LEETCODE_USERNAME")
	password := os.Getenv("LEETCODE_PASSWORD")
	if err := GenLeetCodeTests(username, password); err != nil {
		t.Fatal(err)
	}
}
