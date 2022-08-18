package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/216/D
// https://codeforces.com/problemset/status/216/problem/D
func TestCF216D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
3 1 6 7
4 3 5 2 9
2 8 1
4 3 7 6 4
3 2 5 9
3 6 3 8
3 4 2 9
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF216D)
}
