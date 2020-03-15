package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF740D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 5 1 4 6
1 7
1 1
3 5
3 6
outputCopy
1 0 1 0 0
inputCopy
5
9 7 8 6 5
1 1
2 1
3 1
4 1
outputCopy
4 3 2 1 0`
	testutil.AssertEqualCase(t, rawText, 0, CF740D)
}
