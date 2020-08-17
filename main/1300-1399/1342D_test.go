package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1342D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
1 2 2 3
4 1 1
outputCopy
3
1 2
2 1 3
1 2
inputCopy
6 10
5 8 1 10 8 7
6 6 4 4 3 2 2 2 1 1
outputCopy
2
3 8 5 7
3 10 8 1
inputCopy
5 1
1 1 1 1 1
5
outputCopy
1
5 1 1 1 1 1
inputCopy
5 1
1 1 1 1 1
1
outputCopy
5
1 1
1 1
1 1
1 1
1 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1342D)
}
