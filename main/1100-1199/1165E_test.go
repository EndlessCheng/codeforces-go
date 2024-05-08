package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1165E(t *testing.T) {
	// just copy from website
	rawText := `
5
1 8 7 2 4
9 7 2 9 3
outputCopy
646
inputCopy
1
1000000
1000000
outputCopy
757402647
inputCopy
2
1 3
4 2
outputCopy
20`
	testutil.AssertEqualCase(t, rawText, -1, cf1165E)
}
