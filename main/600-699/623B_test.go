package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/623/B
// https://codeforces.com/problemset/status/623/problem/B
func TestCF623B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 1 4
4 2 3
outputCopy
1
inputCopy
5 3 2
5 17 13 5 6
outputCopy
8
inputCopy
8 3 4
3 7 5 4 3 12 9 4
outputCopy
13`
	testutil.AssertEqualCase(t, rawText, 0, CF623B)
}
