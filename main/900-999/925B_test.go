package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/925/B
// https://codeforces.com/problemset/status/925/problem/B
func TestCF925B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 8 16
3 5 2 9 8 7
outputCopy
Yes
3 2
1 2 6
5 4
inputCopy
4 20 32
21 11 11 12
outputCopy
Yes
1 3
1
2 3 4
inputCopy
4 11 32
5 5 16 16
outputCopy
No
inputCopy
5 12 20
7 8 4 11 9
outputCopy
No`
	testutil.AssertEqualCase(t, rawText, 0, CF925B)
}
