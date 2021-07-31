package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1555/A
// https://codeforces.com/problemset/status/1555/problem/A
func TestCF1555A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
12
15
300
1
9999999999999999
3
outputCopy
30
40
750
15
25000000000000000
15`
	testutil.AssertEqualCase(t, rawText, 0, CF1555A)
}
