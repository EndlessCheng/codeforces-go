package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1255B(t *testing.T) {
	// just copy from website
	rawText := `
3
4 4
1 1 1 1
3 1
1 2 3
3 3
1 2 3
outputCopy
8
1 2
4 3
3 2
4 1
-1
12
3 2
1 2
3 1`
	testutil.AssertEqualCase(t, rawText, 0, Sol1255B)
}
