package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1114C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 9
outputCopy
1
inputCopy
38 11
outputCopy
3
inputCopy
5 2
outputCopy
3
inputCopy
5 10
outputCopy
1
inputCopy
1000000000000000000 97
outputCopy
10416666666666661`
	testutil.AssertEqualCase(t, rawText, -1, CF1114C)
}
