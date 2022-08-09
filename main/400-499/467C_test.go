package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/467/problem/C
// https://codeforces.com/problemset/status/467/problem/C
func TestCF467C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2 1
1 2 3 4 5
outputCopy
9
inputCopy
7 1 3
2 10 7 18 5 33 0
outputCopy
61`
	testutil.AssertEqualCase(t, rawText, 0, CF467C)
}
