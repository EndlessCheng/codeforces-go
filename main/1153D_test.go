package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1153D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 0 1 1 0 1
1 2 2 2 2
outputCopy
1
inputCopy
5
1 0 1 0 1
1 1 1 1
outputCopy
4
inputCopy
8
1 0 0 1 0 1 1 0
1 1 2 2 3 3 3
outputCopy
4
inputCopy
9
1 1 0 0 1 0 1 0 1
1 1 2 2 3 3 4 4
outputCopy
5
inputCopy
10
1 0 1 1 0 1 1 0 1 1
1 1 2 2 3 3 4 4 5
outputCopy
5
inputCopy
40
1 0 0 0 0 0 1 0 1 1 1 0 0 1 0 0 1 0 1 1 1 0 0 0 0 0 1 0 1 1 1 0 0 0 0 0 1 0 1 1
1 1 2 2 3 3 4 5 5 1 1 2 2 3 3 4 5 5 1 1 2 2 3 3 4 5 5 1 1 2 2 3 3 4 5 5 1 1 2 2 3 3 4 5 5 1 1 2 2 3 3 4 5 5 1 1 2 6 3 3 4 5 5 1 1 2 2 3 3 4 5 5
outputCopy
35`
	testutil.AssertEqualCase(t, rawText, 0, CF1153D)
}
