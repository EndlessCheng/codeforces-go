package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1/C
// https://codeforces.com/problemset/status/1/problem/C
func TestCF1C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
0.000000 0.000000
1.000000 1.000000
0.000000 1.000000
outputCopy
1.00000000
inputCopy
71.756151 7.532275
-48.634784 100.159986
91.778633 158.107739
outputCopy
9991.27897663`
	testutil.AssertEqualCase(t, rawText, -1, CF1C)
}
