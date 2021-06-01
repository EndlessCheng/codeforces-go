package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1109/A
// https://codeforces.com/problemset/status/1109/problem/A
func TestCF1109A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 3 4 5
outputCopy
1
inputCopy
6
3 2 2 3 7 6
outputCopy
3
inputCopy
3
42 4 2
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1109A)
}
