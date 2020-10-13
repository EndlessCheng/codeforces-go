package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF380C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
())(())(())(
7
1 1
2 3
1 2
1 12
8 12
5 11
2 10
outputCopy
0
0
2
10
4
6
6`
	testutil.AssertEqualCase(t, rawText, 0, CF380C)
}
