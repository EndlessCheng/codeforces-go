package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol840D(t *testing.T) {
	// just copy from website
	rawText := `
4 2
33 33 66 66
1 3 2
1 4 2
outputCopy
33
-1
inputCopy
5 3
1 2 1 3 2
2 5 3
1 2 3
5 5 2
outputCopy
2
1
2`
	testutil.AssertEqualCase(t, rawText, 0, Sol840D)
}
