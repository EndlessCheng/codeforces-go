package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1707/A
// https://codeforces.com/problemset/status/1707/problem/A
func TestCF1707A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 1
1
2 1
1 2
3 1
1 2 1
4 2
1 4 3 1
5 2
5 1 2 4 3
outputCopy
1
11
110
1110
01111`
	testutil.AssertEqualCase(t, rawText, 0, CF1707A)
}
