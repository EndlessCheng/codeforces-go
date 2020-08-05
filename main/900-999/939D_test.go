package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF939D(t *testing.T) {
	// just copy from website
	rawText := `
3
abb
dad
outputCopy
2
a d
b a
inputCopy
8
drpepper
cocacola
outputCopy
7
l e
e d
d c
c p
p o
o r
r a`
	testutil.AssertEqualCase(t, rawText, 0, CF939D)
}
