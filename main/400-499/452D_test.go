package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/452/D
// https://codeforces.com/problemset/status/452/problem/D
func TestCF452D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1 1 1 5 5 5
outputCopy
15
inputCopy
8 4 3 2 10 5 2
outputCopy
32`
	testutil.AssertEqualCase(t, rawText, 0, CF452D)
}
