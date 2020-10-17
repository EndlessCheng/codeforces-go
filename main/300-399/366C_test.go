package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF366C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
10 8 1
2 7 1
outputCopy
18
inputCopy
5 3
4 4 4 4 4
2 2 2 2 2
outputCopy
-1
inputCopy
1 1
1
1
outputCopy
1
inputCopy
21 8
50 39 28 27 58 46 95 46 50 8 28 94 61 58 57 7 1 38 9 34 12
94 1 77 1 17 40 99 31 26 1 1 1 15 7 6 1 85 3 32 65 78
outputCopy
352`
	testutil.AssertEqualCase(t, rawText, 0, CF366C)
}
