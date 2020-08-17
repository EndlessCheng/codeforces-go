package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1350B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4
5 3 4 6
7
1 4 2 3 6 4 9
5
5 4 3 2 1
1
9
outputCopy
2
3
1
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1350B)
}
