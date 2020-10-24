package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1406C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
5
1 2
1 3
2 4
2 5
6
1 2
1 3
1 4
2 5
2 6
outputCopy
1 2
1 2
1 3
2 3`
	testutil.AssertEqualCase(t, rawText, 0, CF1406C)
}
