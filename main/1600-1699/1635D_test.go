package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1635/D
// https://codeforces.com/problemset/status/1635/problem/D
func TestCF1635D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 4
6 1
outputCopy
9
inputCopy
4 7
20 39 5 200
outputCopy
14
inputCopy
2 200000
48763 1000000000
outputCopy
448201910`
	testutil.AssertEqualCase(t, rawText, 0, CF1635D)
}
