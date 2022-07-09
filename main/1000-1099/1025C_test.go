package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1025/C
// https://codeforces.com/problemset/status/1025/problem/C
func TestCF1025C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
bwwwbwwbw
outputCopy
5
inputCopy
bwwbwwb
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1025C)
}
