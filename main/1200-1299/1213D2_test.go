package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1213/D2
// https://codeforces.com/problemset/status/1213/problem/D2
func TestCF1213D2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3
1 2 2 4 5
outputCopy
1
inputCopy
5 3
1 2 3 4 5
outputCopy
2
inputCopy
5 3
1 2 3 3 3
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1213D2)
}
