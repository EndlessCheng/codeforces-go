package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF191C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2
1 3
2 4
2 5
2
1 4
3 5
outputCopy
2 1 1 1 
inputCopy
5
3 4
4 5
1 4
2 4
3
2 3
1 3
3 5
outputCopy
3 1 1 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF191C)
}
