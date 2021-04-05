package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1503/A
// https://codeforces.com/problemset/status/1503/problem/A
func TestCF1503A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
6
101101
10
1001101101
4
1100
outputCopy
YES
()()()
((()))
YES
()()((()))
(())()()()
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1503A)
}
