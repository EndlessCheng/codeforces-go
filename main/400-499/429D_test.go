package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/429/D
// https://codeforces.com/problemset/status/429/problem/D
func TestCF429D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 0 0 -1
outputCopy
1
inputCopy
2
1 -1
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF429D)
}
