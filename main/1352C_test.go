package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1352C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
3 7
4 12
2 1000000000
7 97
1000000000 1000000000
2 1
outputCopy
10
15
1999999999
113
1000000001
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1352C)
}
