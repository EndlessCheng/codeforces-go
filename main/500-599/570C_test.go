package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF570C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 3
.b..bz....
1 h
3 c
9 f
outputCopy
4
3
1
inputCopy
4 4
.cc.
2 .
3 .
2 a
1 a
outputCopy
1
3
1
1`
	testutil.AssertEqualCase(t, rawText, 0, CF570C)
}
