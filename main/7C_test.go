package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol7C(t *testing.T) {
	// just copy from website
	rawText := `
2 5 3
outputCopy
6 -3`
	testutil.AssertEqual(t, rawText, Sol7C)
}
