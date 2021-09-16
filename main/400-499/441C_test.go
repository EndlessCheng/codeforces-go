package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/441/C
// https://codeforces.com/problemset/status/441/problem/C
func TestCF441C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3 3
outputCopy
3 1 1 1 2 1 3
3 2 1 2 2 2 3
3 3 1 3 2 3 3
inputCopy
2 3 1
outputCopy
6 1 1 1 2 1 3 2 3 2 2 2 1`
	testutil.AssertEqualCase(t, rawText, 0, CF441C)
}
