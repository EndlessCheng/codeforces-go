package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1255D(t *testing.T) {
	// just copy from website
	rawText := `
4
3 5 3
..R..
...R.
....R
6 4 6
R..R
R..R
RRRR
RRRR
R..R
R..R
5 5 4
RRR..
R.R..
RRR..
R..R.
R...R
2 31 62
RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR
RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR
outputCopy
11122
22223
33333
aacc
aBBc
aBBc
CbbA
CbbA
CCAA
11114
22244
32444
33344
33334
abcdefghijklmnopqrstuvwxyzABCDE
FGHIJKLMNOPQRSTUVWXYZ0123456789`
	testutil.AssertEqualCase(t, rawText, 0, Sol1255D)
}
