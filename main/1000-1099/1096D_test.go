package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1096D(t *testing.T) {
	// just copy from website
	rawText := `
6
hhardh
3 2 9 11 7 1
outputCopy
5
inputCopy
8
hhzarwde
3 2 6 9 4 8 7 1
outputCopy
4
inputCopy
6
hhaarr
1 2 3 4 5 6
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, -1, CF1096D)
}
