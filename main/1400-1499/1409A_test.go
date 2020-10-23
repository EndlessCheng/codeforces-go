package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1409A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
5 5
13 42
18 4
1337 420
123456789 1000000000
100500 9000
outputCopy
0
3
2
92
87654322
9150`
	testutil.AssertEqualCase(t, rawText, 0, CF1409A)
}
