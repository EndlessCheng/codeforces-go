package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF931C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
0 2
outputCopy
0
1 1
inputCopy
6
-1 1 1 0 0 -1
outputCopy
2
0 0 0 0 0 0 
inputCopy
3
100 100 101
outputCopy
3
101 100 100 
inputCopy
7
-10 -9 -10 -8 -10 -9 -9
outputCopy
5
-10 -10 -9 -9 -9 -9 -9 `
	testutil.AssertEqualCase(t, rawText, 0, CF931C)
}
/*
0001112
02 => 11
11 => 02
*/
