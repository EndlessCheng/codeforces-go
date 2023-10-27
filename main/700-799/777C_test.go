package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/777/C
// https://codeforces.com/problemset/status/777/problem/C
func TestCF777C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 4
1 2 3 5
3 1 3 2
4 5 2 3
5 5 3 2
4 4 3 4
6
1 1
2 5
4 5
3 5
1 3
1 5
outputCopy
Yes
No
Yes
Yes
Yes
No`
	testutil.AssertEqualCase(t, rawText, 0, CF777C)
}
