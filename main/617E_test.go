package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol617E(t *testing.T) {
	// just copy from website
	rawText := `
6 2 3
1 2 1 1 0 3
1 6
3 5
outputCopy
7
0
inputCopy
5 3 1
1 1 1 1 1
1 5
2 4
1 3
outputCopy
9
4
4
inputCopy
5 3 0
0 0 0 0 0
1 5
2 4
1 3
outputCopy
15
6
6
inputCopy
29 2 0
1 1 1 1 1 0 1 0 1 0 1 1 0 0 1 0 0 0 0 0 1 0 0 0 0 0 0 0 0
1 10
6 9
outputCopy
25
4`
	testutil.AssertEqualCase(t, rawText, 0, Sol617E)
}
