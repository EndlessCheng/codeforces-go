package _00_399

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol320D(t *testing.T) {
	// just copy from website
	rawText := `
100
61 96 25 10 50 71 38 77 76 75 59 100 89 66 6 99 2 13 3 23 91 93 22 92 4 86 90 44 39 31 9 47 28 95 18 54 1 73 94 78 60 20 42 84 97 83 16 81 67 64 74 46 82 5 88 80 14 48 53 79 30 11 62 21 41 70 63 58 51 56 57 17 87 72 27 85 68 49 52 8 12 98 43 37 35 69 55 32 26 40 29 65 19 24 34 33 15 45 36 7
outputCopy
8
inputCopy
10
10 7 4 2 5 8 9 6 3 1
outputCopy
4
inputCopy
10
10 9 7 8 6 5 3 4 2 1
outputCopy
2
inputCopy
6
1 2 3 4 5 6
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 3, Sol320D)
}
