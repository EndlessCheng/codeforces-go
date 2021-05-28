package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1408/E
// https://codeforces.com/problemset/status/1408/problem/E
func TestCF1408E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
1 2 3
4 5
2 1 2
2 1 2
2 1 2
outputCopy
11
inputCopy
7 8
3 6 7 9 10 7 239
8 1 9 7 10 2 6 239
3 2 1 3
2 4 1
3 1 3 7
2 4 3
5 3 4 5 6 7
2 5 7
1 8
outputCopy
66`
	testutil.AssertEqualCase(t, rawText, 0, CF1408E)
}
