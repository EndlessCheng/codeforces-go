package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF920F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 6
6 4 1 10 3 2 4
2 1 7
2 4 5
1 3 5
2 4 4
1 5 7
2 1 7
outputCopy
30
13
4
22`
	testutil.AssertEqualCase(t, rawText, 0, CF920F)
}
