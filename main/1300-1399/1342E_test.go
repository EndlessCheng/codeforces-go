package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1342/E
// https://codeforces.com/problemset/status/1342/problem/E
func TestCF1342E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
outputCopy
6
inputCopy
3 3
outputCopy
0
inputCopy
4 0
outputCopy
24
inputCopy
1337 42
outputCopy
807905441`
	testutil.AssertEqualCase(t, rawText, 0, CF1342E)
}
