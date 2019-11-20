package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1255C(t *testing.T) {
	// just copy from website
	rawText := `
5
4 3 2
2 3 5
4 1 2
outputCopy
1 4 2 3 5`
	testutil.AssertEqualCase(t, rawText, 0, Sol1255C)
}
