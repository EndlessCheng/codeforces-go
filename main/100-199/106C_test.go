package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF106C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 2 2 1
7 3 2 100
12 3 1 10
outputCopy
241
inputCopy
100 1 25 50
15 5 20 10
outputCopy
200`
	testutil.AssertEqualCase(t, rawText, 0, CF106C)
}
