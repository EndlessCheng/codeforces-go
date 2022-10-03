package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/540/D
// https://codeforces.com/problemset/status/540/problem/D
func TestCF540D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2 2
outputCopy
0.333333333333 0.333333333333 0.333333333333
inputCopy
2 1 2
outputCopy
0.150000000000 0.300000000000 0.550000000000
inputCopy
1 1 3
outputCopy
0.057142857143 0.657142857143 0.285714285714`
	testutil.AssertEqualCase(t, rawText, 0, CF540D)
}
