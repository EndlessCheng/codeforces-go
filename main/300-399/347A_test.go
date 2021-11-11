package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/347/A
// https://codeforces.com/problemset/status/347/problem/A
func TestCF347A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
100 -100 50 0 -50
outputCopy
100 -50 0 50 -100 `
	testutil.AssertEqualCase(t, rawText, 0, CF347A)
}
