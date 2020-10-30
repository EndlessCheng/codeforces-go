package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1213E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
ab
bc
outputCopy
YES
acbbac
inputCopy
3
aa
bc
outputCopy
YES
cacbacbab
inputCopy
1
cb
ac
outputCopy
YES
abc`
	testutil.AssertEqualCase(t, rawText, 0, CF1213E)
}
