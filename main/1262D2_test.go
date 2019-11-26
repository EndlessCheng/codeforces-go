package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1262D2(t *testing.T) {
	// just copy from website
	rawText := `
3
10 20 10
6
1 1
2 1
2 2
3 1
3 2
3 3
outputCopy
20
10
20
10
20
10
inputCopy
7
1 2 1 3 1 2 1
9
2 1
2 2
3 1
3 2
3 3
1 1
7 1
7 7
7 4
outputCopy
2
3
2
3
2
3
1
1
3`
	testutil.AssertEqualCase(t, rawText, 0, Sol1262D2)
}
