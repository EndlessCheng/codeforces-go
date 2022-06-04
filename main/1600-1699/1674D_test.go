package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1674/D
// https://codeforces.com/problemset/status/1674/problem/D
func TestCF1674D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
3 1 5 3
3
3 2 1
1
7331
outputCopy
YES
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1674D)
}
