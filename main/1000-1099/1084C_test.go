package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1084/C
// https://codeforces.com/problemset/status/1084/problem/C
func TestCF1084C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
abbaa
outputCopy
5
inputCopy
baaaa
outputCopy
4
inputCopy
agaa
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1084C)
}
