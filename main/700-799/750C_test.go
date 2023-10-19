package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/750/C
// https://codeforces.com/problemset/status/750/problem/C
func TestCF750C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
-7 1
5 2
8 2
outputCopy
1907
inputCopy
2
57 1
22 2
outputCopy
Impossible
inputCopy
1
-5 1
outputCopy
Infinity
inputCopy
4
27 2
13 1
-50 1
8 2
outputCopy
1897
inputCopy
1
-100 2
outputCopy
1799`
	testutil.AssertEqualCase(t, rawText, 0, CF750C)
}
