package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1059D(t *testing.T) {
	// just copy from website
	rawText := `
1
0 1
outputCopy
0.5
inputCopy
3
0 1
0 2
0 -3
outputCopy
-1
inputCopy
2
0 1
1 1
outputCopy
0.625
inputCopy
4
-10000000 1
10000000 1
-10000000 10000000
10000000 10000000
outputCopy
50000000000000.4949989318847656
inputCopy
2
-10000000 1
10000000 1
outputCopy
50000000000000.4949989318847656`
	testutil.AssertEqualCase(t, rawText, 0, Sol1059D)
}
