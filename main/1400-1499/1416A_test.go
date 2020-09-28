package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1416A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
5
1 2 3 4 5
5
4 4 4 4 2
6
1 3 1 5 3 1
outputCopy
-1 -1 3 2 1 
-1 4 4 4 2 
-1 -1 1 1 1 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1416A)
}
