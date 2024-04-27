package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1521/E
// https://codeforces.com/problemset/status/1521/problem/E
func TestCF1521E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
3 4
2 0 0 1
15 4
2 4 8 1
outputCopy
2
4 1
0 1
5
3 0 0 2 2
3 2 3 3 0
0 1 0 4 0
3 0 0 0 0
2 1 3 3 3`
	testutil.AssertEqualCase(t, rawText, 0, CF1521E)
}
