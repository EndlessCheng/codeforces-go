package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1371D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 2
3 8
1 0
4 16
outputCopy
0
10
01
2
111
111
101
0
0
0
1111
1111
1111
1111`
	testutil.AssertEqualCase(t, rawText, 0, CF1371D)
}
