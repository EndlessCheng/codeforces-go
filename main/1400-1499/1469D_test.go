package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1469/D
// https://codeforces.com/problemset/status/1469/problem/D
func TestCF1469D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
3
4
outputCopy
2
3 2
3 2
3
3 4
4 2
4 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1469D)
}
