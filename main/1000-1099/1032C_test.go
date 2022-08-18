package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1032C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 1 4 2 2
outputCopy
1 4 5 4 5 
inputCopy
7
1 5 7 8 10 3 1
outputCopy
1 2 3 4 5 4 3 
inputCopy
19
3 3 7 9 8 8 8 8 7 7 7 7 5 3 3 3 3 8 8
outputCopy
1 3 4 5 4 5 4 5 4 5 4 5 4 3 5 4 3 5 4 
inputCopy
9
5 4 3 2 1 1 2 3 4
outputCopy
5 4 3 2 1 2 3 4 5
inputCopy
9
1 2 3 4 5 5 4 3 2
outputCopy
1 2 3 4 5 4 3 2 1
inputCopy
6
5 5 4 3 2 1
outputCopy
1 5 4 3 2 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1032C)
}
