package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/231/C
// https://codeforces.com/problemset/status/231/problem/C
func TestCF231C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3
6 3 4 0 2
outputCopy
3 4
inputCopy
3 4
5 5 5
outputCopy
3 5
inputCopy
5 3
3 1 2 2 1
outputCopy
4 2`
	testutil.AssertEqualCase(t, rawText, 0, CF231C)
}
