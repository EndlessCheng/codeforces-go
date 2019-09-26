package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol847H(t *testing.T) {
	// just copy from website
	rawText := `
5
1 4 3 2 5
outputCopy
6
inputCopy
5
1 2 2 2 1
outputCopy
1
inputCopy
7
10 20 40 50 70 90 30
outputCopy
0`
	testutil.AssertEqual(t, rawText, Sol847H)
}
