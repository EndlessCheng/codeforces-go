package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF160D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 5
1 2 101
1 3 100
2 3 2
2 4 2
3 4 1
outputCopy
none
any
at least one
at least one
any
inputCopy
3 3
1 2 1
2 3 1
1 3 2
outputCopy
any
any
none
inputCopy
3 3
1 2 1
2 3 1
1 3 1
outputCopy
at least one
at least one
at least one
inputCopy
4 5
1 2 100
1 3 100
2 3 2
2 4 2
3 4 1
outputCopy
at least one
at least one
at least one
at least one
any`
	testutil.AssertEqualCase(t, rawText, -1, CF160D)
}