package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1187D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
7
1 7 1 4 4 5 6
1 1 4 4 5 7 6
5
1 1 3 3 5
1 1 3 3 5
2
1 1
1 2
3
1 2 3
3 2 1
outputCopy
YES
YES
NO
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1187D)
}
