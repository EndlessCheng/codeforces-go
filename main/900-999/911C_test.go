package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/911/C
// https://codeforces.com/problemset/status/911/problem/C
func TestCF911C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2 3
outputCopy
YES
inputCopy
4 2 3
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF911C)
}
