package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol339D(t *testing.T) {
	// just copy from website
	rawText := `
2 4
1 6 3 5
1 4
3 4
1 2
1 2
outputCopy
1
3
3
3`
	testutil.AssertEqual(t, rawText, Sol339D)
}
