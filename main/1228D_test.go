package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1228D(t *testing.T) {
	// just copy from website
	rawText := `
6 11
1 2
1 3
1 4
1 5
1 6
2 4
2 5
2 6
3 4
3 5
3 6
outputCopy
1 2 2 3 3 3 
inputCopy
4 6
1 2
1 3
1 4
2 3
2 4
3 4
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, -1, Sol1228D)
}
