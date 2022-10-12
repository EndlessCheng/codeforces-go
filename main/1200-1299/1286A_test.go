package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1286/A
// https://codeforces.com/problemset/status/1286/problem/A
func TestCF1286A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
0 5 0 2 3
outputCopy
2
inputCopy
7
1 0 0 5 0 0 2
outputCopy
1
inputCopy
20
0 0 0 0 0 0 0 0 0 9 16 19 3 6 11 1 7 4 13 12
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, 0, CF1286A)
}
