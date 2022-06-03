package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/689/D
// https://codeforces.com/problemset/status/689/problem/D
func TestCF689D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 2 3 2 1 4
6 7 1 2 3 2
outputCopy
2
inputCopy
3
3 3 3
1 1 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF689D)
}
