package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1702/F
// https://codeforces.com/problemset/status/1702/problem/F
func TestCF1702F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4
2 4 5 24
1 4 6 11
3
1 4 17
4 5 31
5
4 7 10 13 14
2 14 14 26 42
5
2 2 4 4 4
28 46 62 71 98
6
1 2 10 16 64 80
20 43 60 74 85 99
outputCopy
YES
NO
YES
YES
YES
inputCopy
1
5
2 2 4 4 4
28 46 62 71 98
outputCopy
YES`
	testutil.AssertEqualCase(t, rawText, -1, CF1702F)
}
