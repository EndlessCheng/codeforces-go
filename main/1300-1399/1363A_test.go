package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1363A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 1
999
1 1
1000
2 1
51 50
2 2
51 50
3 3
101 102 103
outputCopy
Yes
No
Yes
Yes
No`
	testutil.AssertEqualCase(t, rawText, 0, CF1363A)
}
