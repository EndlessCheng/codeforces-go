package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF977F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
3 3 4 7 5 6 8
outputCopy
4
2 3 5 6 
inputCopy
6
1 3 5 2 4 6
outputCopy
2
1 4 
inputCopy
4
10 9 8 7
outputCopy
1
1 
inputCopy
9
6 7 8 3 4 5 9 10 11
outputCopy
6
1 2 3 7 8 9 `
	testutil.AssertEqualCase(t, rawText, 0, CF977F)
}
