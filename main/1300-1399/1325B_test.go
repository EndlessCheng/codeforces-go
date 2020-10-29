package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1325B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
3
3 2 1
6
3 1 4 1 5 9
outputCopy
3
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1325B)
}
