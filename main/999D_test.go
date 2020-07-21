package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF999D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 3
3 2 0 6 10 12
outputCopy
3
3 2 0 7 10 14 
inputCopy
4 2
0 1 2 3
outputCopy
0
0 1 2 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF999D)
}
