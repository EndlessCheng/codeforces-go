package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/899/C
// https://codeforces.com/problemset/status/899/problem/C
func TestCF899C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
outputCopy
0
2 1 4 
inputCopy
2
outputCopy
1
1 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF899C)
}
