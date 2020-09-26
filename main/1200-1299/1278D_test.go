package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1278D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
9 12
2 11
1 3
6 10
5 7
4 8
outputCopy
YES
inputCopy
5
1 3
2 4
5 9
6 8
7 10
outputCopy
NO
inputCopy
5
5 8
3 6
2 9
7 10
1 4
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1278D)
}
