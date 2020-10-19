package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1421A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
6 12
4 9
59 832
28 14
4925 2912
1 1
outputCopy
10
13
891
18
6237
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1421A)
}
