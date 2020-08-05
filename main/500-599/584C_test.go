package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF584C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
abc
xyc
outputCopy
ayd
inputCopy
1 0
c
b
outputCopy
-1
inputCopy
4 2
acbb
babc
outputCopy
aaba`
	testutil.AssertEqualCase(t, rawText, 0, CF584C)
}
