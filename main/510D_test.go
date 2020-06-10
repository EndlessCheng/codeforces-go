package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF510D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
100 99 9900
1 1 1
outputCopy
2
inputCopy
5
10 20 30 40 50
1 1 1 1 1
outputCopy
-1
inputCopy
7
15015 10010 6006 4290 2730 2310 1
1 1 1 1 1 1 10
outputCopy
6
inputCopy
8
4264 4921 6321 6984 2316 8432 6120 1026
4264 4921 6321 6984 2316 8432 6120 1026
outputCopy
7237`
	testutil.AssertEqualCase(t, rawText, 0, CF510D)
}
