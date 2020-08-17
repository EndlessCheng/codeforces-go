package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1366D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10
2 3 4 5 6 7 8 9 10 24
outputCopy
-1 -1 -1 -1 3 -1 -1 -1 2 2 
-1 -1 -1 -1 2 -1 -1 -1 5 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF1366D)
}
