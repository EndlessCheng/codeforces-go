package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF519B(t *testing.T) {
	// just copy from website
	rawText := `
5
1 5 8 123 7
123 7 5 1
5 1 7
outputCopy
8
123
inputCopy
6
1 4 3 3 5 7
3 7 5 4 3
4 3 7 5
outputCopy
1
3`
	testutil.AssertEqualCase(t, rawText, 0, CF519B)
}
