package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/449/D
// https://codeforces.com/problemset/status/449/problem/D
func TestCF449D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 3 3
outputCopy
0
inputCopy
4
0 1 2 3
outputCopy
10
inputCopy
6
5 2 0 5 2 1
outputCopy
53`
	testutil.AssertEqualCase(t, rawText, 0, CF449D)
}
