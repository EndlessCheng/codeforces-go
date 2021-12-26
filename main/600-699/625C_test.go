package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/625/C
// https://codeforces.com/problemset/status/625/problem/C
func TestCF625C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 1
outputCopy
28
1 2 3 4
5 6 7 8
9 10 11 12
13 14 15 16
inputCopy
5 3
outputCopy
85
5 6 17 18 19
9 10 23 24 25
7 8 20 21 22
3 4 14 15 16
1 2 11 12 13`
	testutil.AssertEqualCase(t, rawText, 0, CF625C)
}
