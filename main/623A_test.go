package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol623A(t *testing.T) {
	// just copy from website
	rawText := `
4 4
1 2
1 3
1 4
3 4
outputCopy
Yes
bacc
inputCopy
2 1
1 2
outputCopy
Yes
aa
inputCopy
4 3
1 2
1 3
1 4
outputCopy
No`
	testutil.AssertEqualCase(t, rawText, -1, Sol623A)
}
