package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1359F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 -1 -1 1 2
2 3 -3 -2 10
-4 2 1 -2 1
-2 -2 -1 2 4
outputCopy
0.585902082262898
inputCopy
2
-1 1 -1 1 200
1 1 1 5 200
outputCopy
No show :(
inputCopy
6
3 2 1 -1 40
1 -3 -1 -1 59
-1 -3 1 1 5
-3 0 -1 1 57
3 1 -1 -1 3
-2 3 1 -1 1
outputCopy
0.707106781186547`
	testutil.AssertEqualCase(t, rawText, 0, CF1359F)
}
