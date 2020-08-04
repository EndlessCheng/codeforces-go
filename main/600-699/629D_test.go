package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF629D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
100 30
40 10
outputCopy
942477.796077000
inputCopy
4
1 1
9 7
1 4
10 7
outputCopy
3983.539484752`
	testutil.AssertEqualCase(t, rawText, 0, CF629D)
}
