package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/730/J
// https://codeforces.com/problemset/status/730/problem/J
func TestCF730J(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 3 4 3
4 7 6 5
outputCopy
2 6
inputCopy
2
1 1
100 100
outputCopy
1 1
inputCopy
5
10 30 5 6 24
10 41 7 8 24
outputCopy
3 11`
	testutil.AssertEqualCase(t, rawText, 0, CF730J)
}
