package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1312B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1
7
4
1 1 3 5
6
3 2 1 5 6 4
outputCopy
7
1 5 1 3
2 4 6 1 3 5`
	testutil.AssertEqualCase(t, rawText, 0, CF1312B)
}
