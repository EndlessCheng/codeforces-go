package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF915D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4
1 2
2 3
3 2
3 1
outputCopy
YES
inputCopy
5 6
1 2
2 3
3 2
3 1
2 1
4 5
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF915D)
}
