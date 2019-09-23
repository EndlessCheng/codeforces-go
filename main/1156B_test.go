package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1156B(t *testing.T) {
	// just copy from website
	rawText := `4
abcd
gg
codeforces
abaca
outputCopy
bdac
gg
codfoerces
No answer`
	testutil.AssertEqual(t, rawText, Sol1156B)
}
