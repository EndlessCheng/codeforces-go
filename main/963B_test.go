package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF963B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
0 1 2 1 2
outputCopy
YES
1
2
3
5
4
inputCopy
4
0 1 2 3
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF963B)
}
