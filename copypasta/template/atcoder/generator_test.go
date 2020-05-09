package atcoder

import (
	"os"
	"testing"
)

func TestGenAtCoderContestTemplates(t *testing.T) {
	username := os.Getenv("ATCODER_USERNAME")
	password := os.Getenv("ATCODER_PASSWORD")
	if err := GenAtCoderContestTemplates(username, password); err != nil {
		t.Fatal(err)
	}
}

// https://atcoder.jp/contests/abc162/tasks/abc162_e
// https://atcoder.jp/contests/abc162/submissions
func TestGenAtCoderProblemTemplate(t *testing.T) {
	const problemURL = ""
	if err := GenAtCoderProblemTemplate(problemURL); err != nil {
		t.Fatal(err)
	}
}
