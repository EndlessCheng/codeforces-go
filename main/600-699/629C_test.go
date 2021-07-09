package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/629/C
// https://codeforces.com/problemset/status/629/problem/C
func TestCF629C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 1
(
outputCopy
4
inputCopy
4 4
(())
outputCopy
1
inputCopy
4 3
(((
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF629C)
}
