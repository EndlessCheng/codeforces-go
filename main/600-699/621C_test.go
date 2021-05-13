package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/621/C
// https://codeforces.com/problemset/status/621/problem/C
func TestCF621C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
1 2
420 421
420420 420421
outputCopy
4500.0
inputCopy
3 5
1 4
2 3
11 14
outputCopy
0.0`
	testutil.AssertEqualCase(t, rawText, 0, CF621C)
}
