package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol910C(t *testing.T) {
	// just copy from website
	rawText := `
3
ab
de
aj
outputCopy
47
inputCopy
5
abcdef
ghij
bdef
accbd
g
outputCopy
136542
inputCopy
3
aa
jj
aa
outputCopy
44`
	testutil.AssertEqual(t, rawText, Sol910C)
}
