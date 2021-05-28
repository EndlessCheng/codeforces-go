package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1243/D
// https://codeforces.com/problemset/status/1243/problem/D
func TestCF1243D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 11
1 3
1 4
1 5
1 6
2 3
2 4
2 5
2 6
3 4
3 5
3 6
outputCopy
2
inputCopy
3 0
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1243D)
}
