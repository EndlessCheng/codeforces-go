package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1312C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4 100
0 0 0 0
1 2
1
3 4
1 4 1
3 2
0 1 3
3 9
0 59049 810
outputCopy
YES
YES
NO
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1312C)
}
