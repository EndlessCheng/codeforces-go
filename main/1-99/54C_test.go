package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/54/C
// https://codeforces.com/problemset/status/54/problem/C
func TestCF54C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
1 2
50
outputCopy
0.500000000000000
inputCopy
2
1 2
9 11
50
outputCopy
0.833333333333333`
	testutil.AssertEqualCase(t, rawText, 0, CF54C)
}
