package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/712/C
// https://codeforces.com/problemset/status/712/problem/C
func TestCF712C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 3
outputCopy
4
inputCopy
8 5
outputCopy
3
inputCopy
22 4
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF712C)
}
