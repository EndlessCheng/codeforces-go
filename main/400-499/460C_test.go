package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/460/C
// https://codeforces.com/problemset/status/460/problem/C
func TestCF460C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 2 3
2 2 2 2 1 1
outputCopy
2
inputCopy
2 5 1
5 8
outputCopy
9`
	testutil.AssertEqualCase(t, rawText, 0, CF460C)
}
