package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol472D(t *testing.T) {
	// just copy from website
	rawText := `
10
0 1 3 4 2 5 1 3 2 7
1 0 2 3 3 4 2 2 3 6
3 2 0 1 5 2 4 4 5 4
4 3 1 0 6 1 5 5 6 3
2 3 5 6 0 7 3 5 4 9
5 4 2 1 7 0 6 6 7 2
1 2 4 5 3 6 0 4 3 8
3 2 4 5 5 6 4 0 5 8
2 3 5 6 4 7 3 5 0 9
7 6 4 3 9 2 8 8 9 0
outputCopy
YES
inputCopy
3
0 2 7
2 0 9
7 9 0
outputCopy
YES
inputCopy
3
1 2 7
2 0 9
7 9 0
outputCopy
NO
inputCopy
3
0 2 2
7 0 9
7 9 0
outputCopy
NO
inputCopy
3
0 1 1
1 0 1
1 1 0
outputCopy
NO
inputCopy
2
0 0
0 0
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, Sol472D)
}
