package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1244C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
627936103814 4254617095171609 45205 1927
outputCopy
94118284813 15672 533817803329
inputCopy
30 60 3 1
outputCopy
20 0 10
inputCopy
10 51 5 4
outputCopy
-1
inputCopy
20 0 15 5
outputCopy
0 0 20`
	testutil.AssertEqualCase(t, rawText, -1, CF1244C)
}
