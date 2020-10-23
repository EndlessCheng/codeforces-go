package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1409C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 1 49
5 20 50
6 20 50
5 3 8
9 13 22
outputCopy
1 49 
20 40 30 50 10
26 32 20 38 44 50 
8 23 18 13 3 
1 10 13 4 19 22 25 16 7 `
	testutil.AssertEqualCase(t, rawText, 0, CF1409C)
}
