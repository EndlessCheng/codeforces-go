package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1059/C
// https://codeforces.com/problemset/status/1059/problem/C
func TestCF1059C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
outputCopy
1 1 3 
inputCopy
2
outputCopy
1 2 
inputCopy
1
outputCopy
1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1059C)
}
