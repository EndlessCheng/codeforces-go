package atcoder

import (
	"os"
	"strings"
	"testing"
)

// 仅限结束的比赛
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
