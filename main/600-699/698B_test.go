package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF698B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 3 3 4
outputCopy
1
2 3 4 4 
inputCopy
5
3 2 2 5 3
outputCopy
0
3 2 2 5 3 
inputCopy
8
2 3 5 4 1 6 6 7
outputCopy
2
2 3 7 8 1 6 6 7 
inputCopy
7
4 3 2 6 3 5 2
outputCopy
1
4 3 3 6 3 5 2`
	testutil.AssertEqualCase(t, rawText, -1, CF698B)
}
