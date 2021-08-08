package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/934/B
// https://codeforces.com/problemset/status/934/problem/B
func TestCF934B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
outputCopy
462
inputCopy
6
outputCopy
8080`
	testutil.AssertEqualCase(t, rawText, 0, CF934B)
}
