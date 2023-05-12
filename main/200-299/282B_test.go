package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/282/B
// https://codeforces.com/problemset/status/282/problem/B
func TestCF282B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 999
999 1
outputCopy
AG
inputCopy
3
400 600
400 600
400 600
outputCopy
AGA`
	testutil.AssertEqualCase(t, rawText, 0, CF282B)
}
