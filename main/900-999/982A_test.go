package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/982/A
// https://codeforces.com/problemset/status/982/problem/A
func TestCF982A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
101
outputCopy
Yes
inputCopy
4
1011
outputCopy
No
inputCopy
5
10001
outputCopy
No`
	testutil.AssertEqualCase(t, rawText, 0, CF982A)
}
