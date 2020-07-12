package main

import (
	"testing"
)

func TestGenCodeforcesContestTemplates(t *testing.T) {
	const contestID = ""
	const overwrite = false
	if err := GenCodeforcesContestTemplates(contestID, overwrite); err != nil {
		t.Fatal(err)
	}
}

// https://codeforces.com/problemset/problem/1293/C
// https://codeforces.com/problemset/status/1291/problem/D
// https://codeforces.com/gym/102253/problem/C
// https://codeforces.com/gym/102253/status/C
func TestGenCodeforcesProblemTemplates(t *testing.T) {
	const problemURL = "https://codeforces.com/problemset/problem//E"
	if err := GenCodeforcesProblemTemplates(problemURL, true); err != nil {
		t.Fatal(err)
	}
}

func TestGenTemplates(t *testing.T) {
	const (
		problemNum = 3
		rootPath   = "../../misc/nowcoder//"
		overwrite  = false
	)
	if err := GenTemplates(problemNum, rootPath, overwrite); err != nil {
		t.Fatal(err)
	}
}
