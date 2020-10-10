package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1419E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
6
4
30
12
outputCopy
2 3 6 
1
2 4 
0
2 30 6 3 15 5 10 
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1419E)
}
