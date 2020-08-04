package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF833A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
2 4
75 45
8 8
16 16
247 994
1000000000 1000000
outputCopy
Yes
Yes
Yes
No
No
Yes`
	testutil.AssertEqualCase(t, rawText, 0, CF833A)
}
