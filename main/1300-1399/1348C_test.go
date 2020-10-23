package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1348C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
4 2
baba
5 2
baacb
5 3
baacb
5 3
aaaaa
6 4
aaxxzz
7 1
phoenix
outputCopy
ab
abbc
b
aa
x
ehinopx`
	testutil.AssertEqualCase(t, rawText, 0, CF1348C)
}
