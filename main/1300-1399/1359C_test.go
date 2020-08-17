package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1359C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
30 10 20
41 15 30
18 13 18
outputCopy
2
7
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1359C)
}
