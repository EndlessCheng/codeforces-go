package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1560/C
// https://codeforces.com/problemset/status/1560/problem/C
func TestCF1560C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
11
14
5
4
1
2
1000000000
outputCopy
2 4
4 3
1 3
2 1
1 1
1 2
31623 14130`
	testutil.AssertEqualCase(t, rawText, 0, CF1560C)
}
