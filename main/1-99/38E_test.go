package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/38/E
// https://codeforces.com/problemset/status/38/problem/E
func TestCF38E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 3
3 4
1 2
outputCopy
5
inputCopy
4
1 7
3 1
5 10
6 1
outputCopy
11`
	testutil.AssertEqualCase(t, rawText, 0, CF38E)
}
