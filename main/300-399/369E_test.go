package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/369/E
// https://codeforces.com/problemset/status/369/problem/E
func TestCF369E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
1 3
4 5
6 7
3 1 4 7
2 4 5
1 8
outputCopy
3
1
0`
	testutil.AssertEqualCase(t, rawText, 0, CF369E)
}
