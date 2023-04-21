package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1095/F
// https://codeforces.com/problemset/status/1095/problem/F
func TestCF1095F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
1 3 3
2 3 5
2 1 1
outputCopy
5
inputCopy
4 0
1 3 3 7
outputCopy
16
inputCopy
5 4
1 2 3 4 5
1 2 8
1 3 10
1 4 7
1 5 15
outputCopy
18
inputCopy
10 1
2 2 9 7 2 3 7 7 1 1
5 8 4
outputCopy
45`
	testutil.AssertEqualCase(t, rawText, 0, CF1095F)
}
