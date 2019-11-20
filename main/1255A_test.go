package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1255A(t *testing.T) {
	// just copy from website
	rawText := `
3
4 0
5 14
3 9
outputCopy
2
3
2`
	testutil.AssertEqualCase(t, rawText, 0, Sol1255A)
}
