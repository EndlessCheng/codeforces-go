package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1358B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5
1 1 2 2 1
6
2 3 4 5 6 7
6
1 5 4 5 1 9
5
1 2 3 5 6
outputCopy
6
1
6
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1358B)
}
