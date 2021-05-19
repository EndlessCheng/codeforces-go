package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/401/D
// https://codeforces.com/problemset/status/401/problem/D
func TestCF401D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
104 2
outputCopy
3
inputCopy
223 4
outputCopy
1
inputCopy
7067678 8
outputCopy
47`
	testutil.AssertEqualCase(t, rawText, 0, CF401D)
}
