package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_p2586(t *testing.T) {
	customInputs := []string{
		`3 5
1 1 2
2 2
5`,
		`3 3 0 0 0 1000`,
		`4 4 1 10 3
2 1 1000`,
		`4 4 1 1 1
1 1 50`,
		`2 2 1 1 1
1 1 3`,
	}
	customAnswers := []string{
		`The game is going on
5
5 1 3 1 4
4 1 3 0 4
3 1 3 0 3
2 1 3 0 2
1 1 4 0 1`,
		`Game over after 57 seconds
6
56 1 4 0 0
55 1 4 2 1
54 1 4 3 1
53 1 4 3 3
52 1 4 2 3
51 1 4 1 2`,
		`Game over after 579 seconds
6
58 31 46 1 1
40 32 6 0 0
30 32 64 3 1
17 32 74 3 3
13 32 44 2 0
9 32 84 2 4`,
		`The game is going on
6
50 1 2 4 3
49 1 0 3 4
48 1 2 4 2
47 1 1 4 1
46 1 1 1 4
45 1 0 1 0`,
		`The game is going on
3
3 1 2 1 2
2 1 3 0 2
1 1 4 0 1`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 5, p2586)
}
