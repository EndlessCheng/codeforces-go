package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1819/C
// https://codeforces.com/problemset/status/1819/problem/C
func TestCF1819C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2
1 3
3 4
3 5
outputCopy
Yes
4 5 1 2 3 
inputCopy
3
1 2
1 3
outputCopy
Yes
1 2 3
inputCopy
15
1 2
1 3
2 4
2 5
3 6
3 7
4 8
4 9
5 10
5 11
6 12
6 13
7 14
7 15
outputCopy
No`
	testutil.AssertEqualCase(t, rawText, 0, CF1819C)
}
