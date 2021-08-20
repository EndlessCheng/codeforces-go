package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1236/C
// https://codeforces.com/problemset/status/1236/problem/C
func TestCF1236C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
outputCopy
2 8 5
9 3 4
7 6 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1236C)
}
