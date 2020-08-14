package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1183F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
5 6 15 30
4
10 6 30 15
3
3 4 6
outputCopy
30
31
10`
	testutil.AssertEqualCase(t, rawText, 0, CF1183F)
}
