package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/19/B
// https://codeforces.com/problemset/status/19/problem/B
func TestCF19B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 10
0 20
1 5
1 3
outputCopy
8
inputCopy
3
0 1
0 10
0 100
outputCopy
111
inputCopy
2
2 87623264
0 864627704
outputCopy
87623264`
	testutil.AssertEqualCase(t, rawText, 0, CF19B)
}
