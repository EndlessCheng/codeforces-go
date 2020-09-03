package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF388C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 100
2 1 10
outputCopy
101 10
inputCopy
1
9 2 8 6 5 9 4 7 1 3
outputCopy
30 15
inputCopy
3
3 1 3 2
3 5 4 6
2 8 7
outputCopy
18 18
inputCopy
3
3 1000 1000 1000
6 1000 1000 1000 1000 1000 1000
5 1000 1000 1000 1000 1000
outputCopy
7000 7000`
	testutil.AssertEqualCase(t, rawText, 0, CF388C)
}
