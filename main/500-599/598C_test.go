package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF598C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
-1 0
0 -1
1 0
1 1
outputCopy
3 4
inputCopy
6
-1 0
0 -1
1 0
1 1
-4 -5
-4 -6
outputCopy
6 5
inputCopy
4
-6427 -6285
-5386 -5267
-3898 7239
-3905 7252
outputCopy
4 3
inputCopy
4
9800 9981
61 9899
-9926 -9932
-149 -9926
outputCopy
3 4
inputCopy
3
-5 1
-5 -1
5 0
outputCopy
1 2`
	testutil.AssertEqualCase(t, rawText, 0, cf598C)
}
