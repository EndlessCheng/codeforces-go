package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/707/D
// https://codeforces.com/problemset/status/707/problem/D
func TestCF707D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 3 3
1 1 1
3 2
4 0
outputCopy
1
4
0
inputCopy
4 2 6
3 2
2 2 2
3 3
3 2
2 2 2
3 2
outputCopy
2
1
3
3
2
4
inputCopy
2 2 2
3 2
2 2 1
outputCopy
2
1`
	testutil.AssertEqualCase(t, rawText, 0, CF707D)
}
