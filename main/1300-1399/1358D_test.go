package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1358D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
1 3 1
outputCopy
5
inputCopy
3 6
3 3 3
outputCopy
12
inputCopy
5 6
4 2 3 1 3
outputCopy
15`
	testutil.AssertEqualCase(t, rawText, 0, CF1358D)
}
