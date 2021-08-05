package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/545/C
// https://codeforces.com/problemset/status/545/problem/C
func TestCF545C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2
2 1
5 10
10 9
19 1
outputCopy
3
inputCopy
5
1 2
2 1
5 10
10 9
20 1
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF545C)
}
