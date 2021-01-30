package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1399/problem/A
// https://codeforces.com/problemset/status/1399/problem/A
func TestCF1399A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3
1 2 2
4
5 5 5 5
3
1 2 4
4
1 3 4 4
1
100
outputCopy
YES
YES
NO
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1399A)
}
