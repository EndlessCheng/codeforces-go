package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/425/D
// https://codeforces.com/problemset/status/425/problem/D
func TestCF425D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
0 0
0 2
2 0
2 2
1 1
outputCopy
1
inputCopy
9
0 0
1 1
2 2
0 1
1 0
0 2
2 0
1 2
2 1
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF425D)
}
