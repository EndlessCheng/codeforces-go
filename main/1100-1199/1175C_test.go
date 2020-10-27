package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1175C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 2
1 2 5
2 1
1 1000000000
1 0
4
outputCopy
3
500000000
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1175C)
}
