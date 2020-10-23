package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1409D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 1
1 1
500 4
217871987498122 10
100000000000000001 1
outputCopy
8
0
500
2128012501878
899999999999999999`
	testutil.AssertEqualCase(t, rawText, 0, CF1409D)
}
