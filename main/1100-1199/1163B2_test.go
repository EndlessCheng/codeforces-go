package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1163/B2
// https://codeforces.com/problemset/status/1163/problem/B2
func TestCF1163B2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
13
1 1 1 2 2 2 3 3 3 4 4 4 5
outputCopy
13
inputCopy
5
10 100 20 200 1
outputCopy
5
inputCopy
1
100000
outputCopy
1
inputCopy
7
3 2 1 1 4 5 1
outputCopy
6
inputCopy
6
1 1 1 2 2 2
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1163B2)
}
