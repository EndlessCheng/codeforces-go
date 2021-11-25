package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1209/C
// https://codeforces.com/problemset/status/1209/problem/C
func TestCF1209C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
12
040425524644
1
0
9
123456789
2
98
3
987
outputCopy
121212211211
1
222222222
21
-`
	testutil.AssertEqualCase(t, rawText, 0, CF1209C)
}
