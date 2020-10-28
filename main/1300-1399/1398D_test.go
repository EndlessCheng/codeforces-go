package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1398D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1 1
3
5
4
outputCopy
20
inputCopy
2 1 3
9 5
1
2 8 5
outputCopy
99
inputCopy
10 1 1
11 7 20 15 19 14 2 4 13 14
8
11
outputCopy
372`
	testutil.AssertEqualCase(t, rawText, 0, CF1398D)
}
