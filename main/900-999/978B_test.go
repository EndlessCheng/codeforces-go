package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF978B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
xxxiii
outputCopy
1
inputCopy
5
xxoxx
outputCopy
0
inputCopy
10
xxxxxxxxxx
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, 0, CF978B)
}
