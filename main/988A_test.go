package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF988A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3
15 13 15 15 12
outputCopy
YES
1 2 5 
inputCopy
5 4
15 13 15 15 12
outputCopy
NO
inputCopy
4 4
20 10 40 30
outputCopy
YES
1 2 3 4 `
	testutil.AssertEqualCase(t, rawText, 0, CF988A)
}
