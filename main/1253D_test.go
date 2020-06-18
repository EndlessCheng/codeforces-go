package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1253D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
14 8
1 2
2 7
3 4
6 3
5 7
3 8
6 8
11 12
outputCopy
1
inputCopy
200000 3
7 9
9 8
4 5
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1253D)
}
