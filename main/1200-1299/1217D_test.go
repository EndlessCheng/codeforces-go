package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1217D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 5
1 2
1 3
3 4
2 4
1 4
outputCopy
1
1 1 1 1 1 
inputCopy
3 3
1 2
2 3
3 1
outputCopy
2
1 1 2 
inputCopy
6 8
1 2
2 3
3 1
4 3
5 4
6 5
1 6
6 2
outputCopy
2
1 1 2 1 1 1 1 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1217D)
}
