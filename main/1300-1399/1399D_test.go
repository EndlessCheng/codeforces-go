package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1399D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4
0011
6
111111
5
10101
8
01010000
outputCopy
2
1 2 2 1 
6
1 2 3 4 5 6 
1
1 1 1 1 1 
4
1 1 1 1 1 2 3 4 `
	testutil.AssertEqualCase(t, rawText, 0, CF1399D)
}
