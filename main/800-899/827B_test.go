package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF827B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
outputCopy
2
1 2
2 3
inputCopy
5 3
outputCopy
3
1 2
2 3
3 4
3 5
inputCopy
8 3
outputCopy
5
inputCopy
1000 245
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF827B)
}
