package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1307B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2 4
1 3
3 12
3 4 5
1 5
5
2 10
15 4
outputCopy
2
3
1
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1307B)
}
