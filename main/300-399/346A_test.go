package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/346/problem/A
// https://codeforces.com/problemset/status/346/problem/A
func TestCF346A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
2 3
outputCopy
Alice
inputCopy
2
5 3
outputCopy
Alice
inputCopy
3
5 6 7
outputCopy
Bob`
	testutil.AssertEqualCase(t, rawText, 0, CF346A)
}
