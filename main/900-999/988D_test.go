package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF988D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
3 5 4 7 10 12
outputCopy
3
7 3 5
inputCopy
5
-1 2 5 8 11
outputCopy
1
8`
	testutil.AssertEqualCase(t, rawText, 0, CF988D)
}
