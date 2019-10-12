package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol935C(t *testing.T) {
	// just copy from website
	rawText := `
5 3 3 1 1
outputCopy
3.7677669529663684 3.7677669529663684 3.914213562373095
inputCopy
10 5 5 5 15
outputCopy
5.0 5.0 10.0`
	testutil.AssertEqualCase(t, rawText, -1, Sol935C)
}
