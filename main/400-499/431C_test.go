package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol431C(t *testing.T) {
	// just copy from website
	rawText := `
3 3 2
outputCopy
3
inputCopy
3 3 3
outputCopy
1
inputCopy
4 3 2
outputCopy
6
inputCopy
4 5 2
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, 0, Sol431C)
}
