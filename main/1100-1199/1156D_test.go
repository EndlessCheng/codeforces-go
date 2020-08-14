package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1156D(t *testing.T) {
	// just copy from website
	rawText := `
7
2 1 1
3 2 0
4 2 1
5 2 0
6 7 1
7 2 1
outputCopy
34`
	testutil.AssertEqualCase(t, rawText, -1, Sol1156D)
}
