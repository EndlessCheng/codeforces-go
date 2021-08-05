package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/453/A
// https://codeforces.com/problemset/status/453/problem/A
func TestCF453A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 1
outputCopy
3.500000000000
inputCopy
6 3
outputCopy
4.958333333333
inputCopy
2 2
outputCopy
1.750000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF453A)
}
