package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1651/C
// https://codeforces.com/problemset/status/1651/problem/C
func TestCF1651C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
3
1 10 1
20 4 25
4
1 1 1 1
1000000000 1000000000 1000000000 1000000000
outputCopy
31
1999999998
inputCopy
1
3
1 4 3
2 1 4
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, -1, CF1651C)
}
