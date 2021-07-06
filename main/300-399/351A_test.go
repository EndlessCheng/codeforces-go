package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/351/A
// https://codeforces.com/problemset/status/351/problem/A
func TestCF351A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
0.000 0.500 0.750 1.000 2.000 3.000
outputCopy
0.250
inputCopy
3
4469.000 6526.000 4864.000 9356.383 7490.000 995.896
outputCopy
0.279`
	testutil.AssertEqualCase(t, rawText, 0, CF351A)
}
