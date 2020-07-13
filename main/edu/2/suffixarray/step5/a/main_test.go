package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
ababba
outputCopy
15
inputCopy
mmuc
outputCopy
9
inputCopy
xmnnnuu
outputCopy
24
inputCopy
nnnn
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
