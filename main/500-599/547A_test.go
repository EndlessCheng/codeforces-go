package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/547/A
// https://codeforces.com/problemset/status/547/problem/A
func TestCF547A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4 2
1 1
0 1
2 3
outputCopy
3
inputCopy
1023
1 2
1 0
1 2
1 1
outputCopy
-1
inputCopy
29
4 0
1 1
25 20
16 0
outputCopy
170`
	testutil.AssertEqualCase(t, rawText, 0, CF547A)
}
