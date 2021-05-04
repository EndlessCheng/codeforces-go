package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1098/B
// https://codeforces.com/problemset/status/1098/problem/B
func TestCF1098B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2
AG
CT
outputCopy
AG
CT
inputCopy
3 5
AGCAG
AGCAG
AGCAG
outputCopy
TGCAT
CATGC
TGCAT`
	testutil.AssertEqualCase(t, rawText, 0, CF1098B)
}
