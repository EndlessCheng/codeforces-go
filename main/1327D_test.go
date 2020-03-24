package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1327D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
1 3 4 2
1 2 2 3
5
2 3 4 5 1
1 2 3 4 5
8
7 4 5 6 1 8 3 2
5 3 6 4 7 5 8 4
outputCopy
1
5
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1327D)
}
