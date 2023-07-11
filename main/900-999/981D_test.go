package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol981D(t *testing.T) {
	// just copy from website
	rawText := `10 4
9 14 28 1 7 13 15 29 2 31
outputCopy
24
inputCopy
7 3
3 14 15 92 65 35 89
outputCopy
64`
	testutil.AssertEqual(t, rawText, CF981D)
}
