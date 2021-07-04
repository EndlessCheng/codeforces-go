package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1296/E2
// https://codeforces.com/problemset/status/1296/problem/E2
func TestCF1296E2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9
abacbecfd
outputCopy
2
1 1 2 1 2 1 2 1 2 
inputCopy
8
aaabbcbb
outputCopy
2
1 2 1 2 1 2 1 1
inputCopy
7
abcdedc
outputCopy
3
1 1 1 1 1 2 3 
inputCopy
5
abcde
outputCopy
1
1 1 1 1 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1296E2)
}
