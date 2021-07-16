package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/319/C
// https://codeforces.com/problemset/status/319/problem/C
func TestCF319C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 3 4 5
5 4 3 2 0
outputCopy
25
inputCopy
6
1 2 3 10 20 30
6 5 4 3 2 0
outputCopy
138`
	testutil.AssertEqualCase(t, rawText, 0, CF319C)
}
