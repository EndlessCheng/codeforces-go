package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1060D(t *testing.T) {
	// just copy from website
	rawText := `
3
1 1
1 1
1 1
outputCopy
6
inputCopy
4
1 2
2 1
3 5
5 3
outputCopy
15
inputCopy
1
5 6
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1060D)
}
