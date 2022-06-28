package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/439/C
// https://codeforces.com/problemset/status/439/problem/C
func TestCF439C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5 3
2 6 10 5 9
outputCopy
YES
1 9
1 5
1 10
1 6
1 2
inputCopy
5 5 3
7 14 2 9 5
outputCopy
NO
inputCopy
5 3 1
1 2 3 7 5
outputCopy
YES
3 5 1 3
1 7
1 2`
	testutil.AssertEqualCase(t, rawText, 0, CF439C)
}
