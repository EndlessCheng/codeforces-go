package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF246D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 6
1 1 2 3 5 8
1 2
3 2
1 4
4 3
4 5
4 6
outputCopy
3
inputCopy
5 6
4 2 5 2 4
1 2
2 3
3 1
5 3
5 4
3 4
outputCopy
2
inputCopy
3 1
13 13 4
1 2
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF246D)
}
