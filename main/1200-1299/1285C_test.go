package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1285/C
// https://codeforces.com/problemset/status/1285/problem/C
func TestCF1285C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
outputCopy
1 2
inputCopy
6
outputCopy
2 3
inputCopy
4
outputCopy
1 4
inputCopy
1
outputCopy
1 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1285C)
}
