package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
bababb
zabacabba
outputCopy
aba
inputCopy
qrdq
rqqqrdqrqd
outputCopy
qrdq
inputCopy
hhhhhh
hhhhhhh
outputCopy
hhhhhh
inputCopy
opopo
ppppopopo
outputCopy
opopo`
	testutil.AssertEqualCase(t, rawText, 0, run)
}
