package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1278C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
6
1 1 1 2 2 1 2 1 2 1 1 2
2
1 2 1 2
3
1 1 1 1 1 1
2
2 1 1 1
outputCopy
6
0
6
2
inputCopy
9
3
1 2 1 2 1 2
1
1 1
1 
2 2
1 
1 2
4
1 1 1 1 2 2 2 2
4
1 1 1 2 2 2 2 2
4
1 1 1 1 1 2 2 2
4
1 1 1 1 1 1 2 2
2
1 2 1 1
outputCopy
0
2
2
0
0
2
2
4
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1278C)
}
