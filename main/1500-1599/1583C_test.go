package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1583/C
// https://codeforces.com/problemset/status/1583/problem/C
func TestCF1583C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 5
..XXX
...X.
...X.
...X.
5
1 3
3 3
4 5
5 5
1 5
outputCopy
YES
YES
NO
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1583C)
}
