package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1385B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2
1 1 2 2
4
1 3 1 4 3 4 2 2
5
1 2 1 2 3 4 3 5 4 5
3
1 2 3 1 2 3
4
2 3 2 4 1 3 4 1
outputCopy
1 2 
1 3 4 2 
1 2 3 4 5 
1 2 3 
2 3 4 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1385B)
}
