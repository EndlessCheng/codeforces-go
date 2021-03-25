package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1506/G
// https://codeforces.com/problemset/status/1506/problem/G
func TestCF1506G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
codeforces
aezakmi
abacaba
convexhull
swflldjgpaxs
myneeocktxpqjpz
outputCopy
odfrces
ezakmi
cba
convexhul
wfldjgpaxs
myneocktxqjpz`
	testutil.AssertEqualCase(t, rawText, 0, CF1506G)
}
