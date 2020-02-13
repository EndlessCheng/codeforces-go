package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1303B(t *testing.T) {
	// just copy from website
	rawText := `
3
5 1 1
8 10 10
1000000 1 1000000
outputCopy
5
8
499999500000`
	testutil.AssertEqualCase(t, rawText, 0, CF1303B)
}
