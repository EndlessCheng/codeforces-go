package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/732/D
// https://codeforces.com/problemset/status/732/problem/D
func TestCF732D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 2
0 1 0 2 1 0 2
2 1
outputCopy
5
inputCopy
10 3
0 0 1 2 3 0 2 0 1 2
1 1 4
outputCopy
9
inputCopy
5 1
1 1 1 1 1
5
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF732D)
}
