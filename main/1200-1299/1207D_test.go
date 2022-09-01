package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1207/D
// https://codeforces.com/problemset/status/1207/problem/D
func TestCF1207D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 1
2 2
3 1
outputCopy
3
inputCopy
4
2 3
2 2
2 1
2 4
outputCopy
0
inputCopy
3
1 1
1 1
2 3
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1207D)
}
