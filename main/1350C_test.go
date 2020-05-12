package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1350C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 1
outputCopy
1
inputCopy
4
10 24 40 80
outputCopy
40
inputCopy
10
540 648 810 648 720 540 594 864 972 648
outputCopy
54`
	testutil.AssertEqualCase(t, rawText, 0, CF1350C)
}
