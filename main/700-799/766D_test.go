package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF766D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3 4
hate love like
1 love like
2 love hate
1 hate like
love like
love hate
like hate
hate like
outputCopy
YES
YES
NO
1
2
2
2
inputCopy
8 6 5
hi welcome hello ihateyou goaway dog cat rat
1 hi welcome
1 ihateyou goaway
2 hello ihateyou
2 hi goaway
2 hi hello
1 hi hello
dog cat
dog hi
hi hello
ihateyou goaway
welcome ihateyou
outputCopy
YES
YES
YES
YES
NO
YES
3
3
1
1
2`
	testutil.AssertEqualCase(t, rawText, 0, CF766D)
}