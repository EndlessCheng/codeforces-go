package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF864E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 7 4
2 6 5
3 7 6
outputCopy
11
2
2 3 
inputCopy
2
5 6 1
3 3 5
outputCopy
1
1
1 
inputCopy
9
13 18 14
8 59 20
9 51 2
18 32 15
1 70 18
14 81 14
10 88 16
18 52 3
1 50 6
outputCopy
106
8
1 4 9 8 2 5 6 7 `
	testutil.AssertEqualCase(t, rawText, 0, CF864E)
}
