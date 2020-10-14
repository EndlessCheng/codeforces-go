package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF86D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
1 2 1
1 2
1 3
outputCopy
3
6
inputCopy
8 3
1 1 2 2 1 3 1 1
2 7
1 6
2 7
outputCopy
20
20
20`
	testutil.AssertEqualCase(t, rawText, 0, CF86D)
}
