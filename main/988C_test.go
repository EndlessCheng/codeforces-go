package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF988C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
5
2 3 1 3 2
6
1 1 2 2 2 1
outputCopy
YES
2 6
1 2
inputCopy
3
1
5
5
1 1 1 1 1
2
2 3
outputCopy
NO
inputCopy
4
6
2 2 2 2 2 2
5
2 2 2 2 2
3
2 2 2
5
2 2 2 2 2
outputCopy
YES
2 2
4 1`
	testutil.AssertEqualCase(t, rawText, 0, CF988C)
}
