package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/559/A
// https://codeforces.com/problemset/status/559/problem/A
func TestCF559A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1 1 1 1 1
outputCopy
6
inputCopy
1 2 1 2 1 2
outputCopy
13`
	testutil.AssertEqualCase(t, rawText, 0, CF559A)
}
