package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/610/B
// https://codeforces.com/problemset/status/610/problem/B
func TestCF610B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 4 2 3 3
outputCopy
12
inputCopy
3
5 5 5
outputCopy
15
inputCopy
6
10 10 10 1 10 10
outputCopy
11`
	testutil.AssertEqualCase(t, rawText, 0, CF610B)
}
