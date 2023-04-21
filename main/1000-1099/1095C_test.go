package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1095/C
// https://codeforces.com/problemset/status/1095/problem/C
func TestCF1095C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9 4
outputCopy
YES
1 2 2 4 
inputCopy
8 1
outputCopy
YES
8 
inputCopy
5 1
outputCopy
NO
inputCopy
3 7
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1095C)
}
