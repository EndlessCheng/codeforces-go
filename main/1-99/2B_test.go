package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF2B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 3
4 5 6
7 8 9
outputCopy
0
DDRR`
	testutil.AssertEqualCase(t, rawText, 0, CF2B)
}
