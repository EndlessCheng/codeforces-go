package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol567D(t *testing.T) {
	// just copy from website
	rawText := `
5000 1660 2
20
1 100 18 102 300 81 19 25 44 88 1337 4999 1054 1203 91 16 164 914 1419 1487
outputCopy
18
inputCopy
11 3 3
5
4 8 6 1 11
outputCopy
3
inputCopy
5 1 3
2
1 5
outputCopy
-1
inputCopy
5 1 3
1
3
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, -1, Sol567D)
}
