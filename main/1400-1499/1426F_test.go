package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1426F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
ac?b?c
outputCopy
24
inputCopy
7
???????
outputCopy
2835
inputCopy
9
cccbbbaaa
outputCopy
0
inputCopy
5
a???c
outputCopy
46`
	testutil.AssertEqualCase(t, rawText, 0, CF1426F)
}
