package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/859/E
// https://codeforces.com/problemset/status/859/problem/E
func TestCF859E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 5
5 2
3 7
7 3
outputCopy
6
inputCopy
5
1 10
2 10
3 10
4 10
5 5
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF859E)
}
