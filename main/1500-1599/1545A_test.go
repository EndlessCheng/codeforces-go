package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1545/A
// https://codeforces.com/problemset/status/1545/problem/A
func TestCF1545A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
4 3 2 5
4
3 3 2 2
5
1 2 3 5 4
outputCopy
YES
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1545A)
}
