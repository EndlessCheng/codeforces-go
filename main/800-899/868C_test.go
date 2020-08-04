package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF868C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3
1 0 1
1 1 0
1 0 0
1 0 0
1 0 0
outputCopy
NO
inputCopy
3 2
1 0
1 1
0 1
outputCopy
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF868C)
}
