package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/768/D
// https://codeforces.com/problemset/status/768/problem/D
func TestCF768D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1
1
outputCopy
1
inputCopy
2 2
1
2
outputCopy
2
2`
	testutil.AssertEqualCase(t, rawText, 0, CF768D)
}
