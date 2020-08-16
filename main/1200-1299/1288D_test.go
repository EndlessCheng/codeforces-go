package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1288D(t *testing.T) {
	// just copy from website
	rawText := `
6 5
5 0 3 1 2
1 8 9 1 3
1 2 3 4 5
9 1 0 3 7
2 3 0 6 3
6 4 1 7 0
outputCopy
1 5`
	testutil.AssertEqualCase(t, rawText, 0, CF1288D)
}
