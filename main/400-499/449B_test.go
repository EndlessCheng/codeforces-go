package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF449B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5 3
1 2 1
2 3 2
1 3 3
3 4 4
1 5 5
3 5
4 5
5 5
outputCopy
2
inputCopy
2 2 3
1 2 2
2 1 3
2 1
2 2
2 3
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF449B)
}
