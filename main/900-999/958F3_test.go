package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol958F3(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3 2
1 2 3 2
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF958F3)
}
