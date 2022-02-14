package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/263/A
// https://codeforces.com/problemset/status/263/problem/A
func TestCF263A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
0 0 0 0 0
0 0 0 0 1
0 0 0 0 0
0 0 0 0 0
0 0 0 0 0
outputCopy
3
inputCopy
0 0 0 0 0
0 0 0 0 0
0 1 0 0 0
0 0 0 0 0
0 0 0 0 0
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF263A)
}
