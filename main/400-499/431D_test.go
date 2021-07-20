package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/431/D
// https://codeforces.com/problemset/status/431/problem/D
func TestCF431D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1
outputCopy
1
inputCopy
3 2
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF431D)
}
