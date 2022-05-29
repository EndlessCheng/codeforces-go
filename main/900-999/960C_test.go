package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/960/C
// https://codeforces.com/problemset/status/960/problem/C
func TestCF960C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 5
outputCopy
6
5 50 7 15 6 100
inputCopy
4 2
outputCopy
4
10 100 1000 10000`
	testutil.AssertEqualCase(t, rawText, 0, CF960C)
}
