package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol242E(t *testing.T) {
	// just copy from website
	rawText := `
5
4 10 3 13 7
8
1 2 4
2 1 3 3
1 2 4
1 3 3
2 2 5 5
1 1 5
2 1 2 10
1 2 3
outputCopy
26
22
0
34
11
inputCopy
6
4 7 4 0 7 3
5
2 2 3 8
1 1 5
2 3 5 1
2 4 5 6
1 2 3
outputCopy
38
28`
	testutil.AssertEqualCase(t, rawText, -1, Sol242E)
}
