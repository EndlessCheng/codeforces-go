package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1105/C
// https://codeforces.com/problemset/status/1105/problem/C
func TestCF1105C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 1 3
outputCopy
3
inputCopy
3 2 2
outputCopy
1
inputCopy
9 9 99
outputCopy
711426616`
	testutil.AssertEqualCase(t, rawText, 0, CF1105C)
}
