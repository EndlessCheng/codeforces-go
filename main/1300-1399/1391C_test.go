package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1391/C
// https://codeforces.com/problemset/status/1391/problem/C
func TestCF1391C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
outputCopy
16
inputCopy
583291
outputCopy
135712853`
	testutil.AssertEqualCase(t, rawText, 0, CF1391C)
}
