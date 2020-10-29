package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1187C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 4
1 1 3
1 2 5
0 5 6
1 6 7
outputCopy
YES
1 2 2 3 5 4 4
inputCopy
4 2
1 1 4
0 2 3
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1187C)
}
