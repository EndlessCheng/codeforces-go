package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1016D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 3
2 9
5 3 13
outputCopy
YES
3 4 5
6 7 8
inputCopy
3 3
1 7 6
2 15 12
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1016D)
}
