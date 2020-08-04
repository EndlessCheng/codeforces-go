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
4 3`
	testutil.AssertEqualCase(t, rawText, 0, CF598C)
}
