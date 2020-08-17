package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1328E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 6
1 2
1 3
1 4
2 5
2 6
3 7
7 8
7 9
9 10
4 3 8 9 10
3 2 4 6
3 2 1 5
3 4 8 2
2 6 10
3 5 4 7
outputCopy
YES
YES
YES
YES
NO
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1328E)
}
