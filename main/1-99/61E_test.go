package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/61/E
// https://codeforces.com/problemset/status/61/problem/E
func TestCF61E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 2 1
outputCopy
1
inputCopy
3
2 3 1
outputCopy
0
inputCopy
4
10 8 3 1
outputCopy
4
inputCopy
4
1 5 4 3
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF61E)
}
