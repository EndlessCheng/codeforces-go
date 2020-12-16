package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF474E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
1 3 6 7 4
outputCopy
4
1 2 3 5 
inputCopy
10 3
2 1 3 6 9 11 7 3 20 18
outputCopy
6
1 4 6 7 8 9 `
	testutil.AssertEqualCase(t, rawText, 0, CF474E)
}
