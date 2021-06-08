package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/884/D
// https://codeforces.com/problemset/status/884/problem/D
func TestCF884D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 3
outputCopy
6
inputCopy
4
2 3 4 5
outputCopy
19`
	testutil.AssertEqualCase(t, rawText, 0, CF884D)
}
