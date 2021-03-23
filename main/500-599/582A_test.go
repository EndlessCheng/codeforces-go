package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/582/A
// https://codeforces.com/problemset/status/582/problem/A
func TestCF582A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 1 2 3 4 3 2 6 1 1 2 2 1 2 3 2
outputCopy
4 3 6 2
inputCopy
1
42
outputCopy
42 
inputCopy
2
1 1 1 1
outputCopy
1 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF582A)
}
