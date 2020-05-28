package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF519E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2
1 3
2 4
1
2 3
outputCopy
1
inputCopy
4
1 2
2 3
2 4
2
1 2
1 3
outputCopy
0
2`
	testutil.AssertEqualCase(t, rawText, 0, CF519E)
}
