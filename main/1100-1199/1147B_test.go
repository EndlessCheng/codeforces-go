package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1147B(t *testing.T) {
	// just copy from website
	rawText := `
12 6
1 3
3 7
5 7
7 11
9 11
11 3
outputCopy
Yes
inputCopy
9 6
4 5
5 6
7 8
8 9
1 2
2 3
outputCopy
Yes
inputCopy
10 3
1 2
3 2
7 2
outputCopy
No
inputCopy
10 2
1 6
2 7
outputCopy
Yes`
	testutil.AssertEqual(t, rawText, Sol1147B)
}
