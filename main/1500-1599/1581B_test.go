package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1581/B
// https://codeforces.com/problemset/status/1581/problem/B
func TestCF1581B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 0 3
4 5 3
4 6 3
5 4 1
2 1 1
outputCopy
YES
NO
YES
NO
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1581B)
}
