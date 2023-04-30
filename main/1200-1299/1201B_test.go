package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1201/B
// https://codeforces.com/problemset/status/1201/problem/B
func TestCF1201B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1 2 2
outputCopy
YES
inputCopy
6
1 2 3 4 5 6
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1201B)
}
