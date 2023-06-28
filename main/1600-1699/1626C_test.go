package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1626/C
// https://codeforces.com/problemset/status/1626/problem/C
func TestCF1626C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1
6
4
2
4 5
2 2
3
5 7 9
2 1 2
outputCopy
10
6
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1626C)
}
