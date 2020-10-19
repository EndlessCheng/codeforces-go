package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1421D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
-3 1
1 3 5 7 9 11
1000000000 1000000000
1000000000 1000000000 1000000000 1000000000 1000000000 1000000000
outputCopy
18
1000000000000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF1421D)
}
