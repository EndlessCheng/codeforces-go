package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol279C(t *testing.T) {
	// just copy from website
	rawText := `
8 6
1 2 2 2 3 2 2 2
1 3
2 3
2 4
8 8
1 4
5 8
outputCopy
Yes
Yes
No
Yes
No
Yes`
	testutil.AssertEqualCase(t, rawText, 0, Sol279C)
}
