package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1221D(t *testing.T) {
	// just copy from website
	rawText := `
3
3
2 4
2 1
3 5
3
2 3
2 10
2 6
4
1 7
3 3
2 6
1000000000 2
outputCopy
2
9
0`
	testutil.AssertEqual(t, rawText, Sol1221D)
}
