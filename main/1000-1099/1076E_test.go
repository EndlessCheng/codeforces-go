package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1076E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2
1 3
2 4
2 5
3
1 1 1
2 0 10
4 10 100
outputCopy
1 11 1 100 0 
inputCopy
5
2 3
2 1
5 4
3 4
5
2 0 4
3 10 1
1 2 3
2 3 10
1 1 7
outputCopy
10 24 14 11 11 `
	testutil.AssertEqualCase(t, rawText, 0, CF1076E)
}
