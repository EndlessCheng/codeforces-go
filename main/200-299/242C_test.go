package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF242C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 7 6 11
3
5 3 8
6 7 11
5 2 5
outputCopy
4
inputCopy
3 4 3 10
3
3 1 4
4 5 9
3 10 10
outputCopy
6
inputCopy
1 1 2 10
2
1 1 3
2 6 10
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF242C)
}
