package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/855/E
// https://codeforces.com/problemset/status/855/problem/E
func TestCF855E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
2 4 9
3 1 10
outputCopy
1
2
inputCopy
2
2 1 100
5 1 100
outputCopy
21
4
inputCopy
1
10 1 99
outputCopy
9`
	testutil.AssertEqualCase(t, rawText, -1, CF855E)
}
