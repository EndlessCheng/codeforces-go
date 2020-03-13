package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF916C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
outputCopy
7 7
1 2 3
2 3 2
3 4 2
2 4 4
inputCopy
5 4
outputCopy
7 13
1 2 2
1 3 4
1 4 3
4 5 4`
	testutil.AssertEqualCase(t, rawText, 0, CF916C)
}
