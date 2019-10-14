package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1244D(t *testing.T) {
	// just copy from website
	rawText := `
3
3 2 3
4 3 2
3 1 3
1 3
2 3
outputCopy
6
1 3 2
inputCopy
3
3 2 3
4 3 2
3 1 3
1 2
2 3
outputCopy
6
1 3 2
inputCopy
5
3 4 2 1 2
4 2 1 5 4
5 3 2 1 1
1 2
3 2
4 3
5 3
outputCopy
-1
inputCopy
5
3 4 2 1 2
4 2 1 5 4
5 3 2 1 1
1 2
3 2
4 3
5 4
outputCopy
9
1 3 2 1 3`
	testutil.AssertEqualCase(t, rawText, -1, Sol1244D)
}
