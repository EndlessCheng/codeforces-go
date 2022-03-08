package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/281/A
// https://codeforces.com/problemset/status/281/problem/A
func TestCF281A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
ApPLe
outputCopy
ApPLe
inputCopy
konjac
outputCopy
Konjac`
	testutil.AssertEqualCase(t, rawText, 0, CF281A)
}
