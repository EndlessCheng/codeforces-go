package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1396B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1
2
2
1 1
outputCopy
T
HL`
	testutil.AssertEqualCase(t, rawText, 0, CF1396B)
}
