package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF570D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 5
1 1 1 3 3
zacccd
1 1
3 3
4 1
6 1
1 2
outputCopy
Yes
No
Yes
Yes
Yes`
	testutil.AssertEqualCase(t, rawText, 0, CF570D)
}
