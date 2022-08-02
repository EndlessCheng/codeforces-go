package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1043/D
// https://codeforces.com/problemset/status/1043/problem/D
func TestCF1043D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
1 2 3
2 3 1
outputCopy
4
inputCopy
5 6
1 2 3 4 5
2 3 1 4 5
3 4 5 1 2
3 5 4 2 1
2 3 5 4 1
1 2 3 4 5
outputCopy
5
inputCopy
2 2
1 2
2 1
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1043D)
}
