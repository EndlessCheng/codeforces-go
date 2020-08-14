package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1199B(t *testing.T) {
	// just copy from website
	rawText := `
1 2
outputCopy
1.5000000000000
inputCopy
3 5
outputCopy
2.6666666666667`
	testutil.AssertEqualCase(t, rawText, 0, Sol1199B)
}
