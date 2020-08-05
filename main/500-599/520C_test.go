package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol520C(t *testing.T) {
	// just copy from website
	rawText := `
1
C
outputCopy
1
inputCopy
2
AG
outputCopy
4
inputCopy
3
TTT
outputCopy
1`
	testutil.AssertEqual(t, rawText, Sol520C)
}
