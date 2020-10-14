package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1132C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 5
1 4
4 5
5 6
6 7
3 5
outputCopy
7
inputCopy
4 3
1 1
2 2
3 4
outputCopy
2
inputCopy
4 4
1 1
2 2
2 3
3 4
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1132C)
}
