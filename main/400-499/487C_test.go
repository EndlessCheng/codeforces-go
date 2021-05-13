package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/487/C
// https://codeforces.com/problemset/status/487/problem/C
func TestCF487C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
outputCopy
YES
1
4
3
6
5
2
7
inputCopy
6
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF487C)
}
