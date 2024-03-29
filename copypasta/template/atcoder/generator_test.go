package atcoder

import (
	"os"
	"strings"
	"testing"
)

// 仅限结束的比赛
// https://atcoder.jp/contests/abc161/tasks/abc161_f
// https://atcoder.jp/contests/abc161/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc161_f&orderBy=source_length
func TestGenAtCoderProblemTemplate(t *testing.T) {
	raw, err := os.ReadFile("data-atcoder.txt")
	if err != nil {
		t.Fatal(err)
	}
	problemURL := strings.TrimSpace(string(raw))
	if err := GenAtCoderProblemTemplate(problemURL); err != nil {
		t.Fatal(err)
	}
}
