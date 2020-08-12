package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1077D(t *testing.T) {
	// just copy from website
	rawText := `
10 3
1 3 1 3 10 3 7 7 12 3
outputCopy
1 3 3 
inputCopy
7 3
1 2 3 2 4 3 1
outputCopy
1 2 3 
inputCopy
10 4
1 3 1 3 10 3 7 7 12 3
outputCopy
1 3 3 7 
inputCopy
15 2
1 2 1 1 1 2 1 1 2 1 2 1 1 1 1
outputCopy
1 1 `
	testutil.AssertEqual(t, rawText, Sol1077D)
}
