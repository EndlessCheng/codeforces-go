package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/925/C
// https://codeforces.com/problemset/status/925/problem/C
func TestCF925C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 3
outputCopy
No
inputCopy
6
4 7 7 12 31 61
outputCopy
Yes
4 12 7 31 7 61 `
	testutil.AssertEqualCase(t, rawText, 0, CF925C)
}
