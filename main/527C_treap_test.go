package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol527CTreap(t *testing.T) {
	// just copy from website
	rawText := `
4 3 4
H 2
V 2
V 3
V 1
outputCopy
8
4
4
2
inputCopy
7 6 5
H 4
V 3
V 5
H 2
V 1
outputCopy
28
16
12
6
4`
	testutil.AssertEqualCase(t, rawText, 0, Sol527CTreap)
}
