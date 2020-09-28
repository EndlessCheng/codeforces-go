package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1417B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
6 7
1 2 3 4 5 6
3 6
3 3 3
outputCopy
1 0 0 1 1 0 
1 0 0`
	testutil.AssertEqualCase(t, rawText, 0, CF1417B)
}
