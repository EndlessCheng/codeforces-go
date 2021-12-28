package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1060/E
// https://codeforces.com/problemset/status/1060/problem/E
func TestCF1060E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2
1 3
1 4
outputCopy
6
inputCopy
4
1 2
2 3
3 4
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1060E)
}
