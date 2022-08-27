package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/79/C
// https://codeforces.com/problemset/status/79/problem/C
func TestCF79C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
Go_straight_along_this_street
5
str
long
tree
biginteger
ellipse
outputCopy
12 4
inputCopy
IhaveNoIdea
9
I
h
a
v
e
N
o
I
d
outputCopy
0 0
inputCopy
unagioisii
2
ioi
unagi
outputCopy
5 5`
	testutil.AssertEqualCase(t, rawText, 0, CF79C)
}
