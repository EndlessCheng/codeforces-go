package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1136/C
// https://codeforces.com/problemset/status/1136/problem/C
func TestCF1136C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2
1 1
6 1
1 6
1 1
outputCopy
YES
inputCopy
2 2
4 4
4 5
5 4
4 4
outputCopy
NO
inputCopy
3 3
1 2 3
4 5 6
7 8 9
1 4 7
2 5 6
3 8 9
outputCopy
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1136C)
}
