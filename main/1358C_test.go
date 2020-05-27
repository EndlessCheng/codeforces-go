package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1358C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1 2 2
1 2 2 4
179 1 179 100000
5 7 5 7
outputCopy
2
3
1
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1358C)
}
