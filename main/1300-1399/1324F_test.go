package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1324F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9
0 1 1 1 0 0 0 0 1
1 2
1 3
3 4
3 5
2 6
4 7
6 8
5 9
outputCopy
2 2 2 2 2 1 1 0 2 
inputCopy
4
0 0 1 0
1 2
1 3
1 4
outputCopy
0 -1 1 -1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1324F)
}
