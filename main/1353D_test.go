package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1353D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1
2
3
4
5
6
outputCopy
1 
1 2 
2 1 3 
3 1 2 4 
2 4 1 3 5 
3 4 1 5 2 6 `
	testutil.AssertEqualCase(t, rawText, 0, CF1353D)
}
