package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1195/C
// https://codeforces.com/problemset/status/1195/problem/C
func TestCF1195C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
9 3 5 7 3
5 8 1 4 5
outputCopy
29
inputCopy
3
1 2 9
10 1 1
outputCopy
19
inputCopy
1
7
4
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1195C)
}
