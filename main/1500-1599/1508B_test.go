package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1508/B
// https://codeforces.com/problemset/status/1508/problem/B
func TestCF1508B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 1
1 2
3 3
6 5
3 4
outputCopy
1 
-1
2 1 3 
1 2 4 3 5 6 
3 2 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1508B)

	tc := [][2]string{}

	s := "32\n"
	for i := 1; i <= 32; i++ {
		s += fmt.Sprintln(6, i)
	}
	tc = append(tc, [2]string{
		s,
		`1 2 3 4 5 6
1 2 3 4 6 5
1 2 3 5 4 6
1 2 3 6 5 4
1 2 4 3 5 6
1 2 4 3 6 5
1 2 5 4 3 6
1 2 6 5 4 3
1 3 2 4 5 6
1 3 2 4 6 5
1 3 2 5 4 6
1 3 2 6 5 4
1 4 3 2 5 6
1 4 3 2 6 5
1 5 4 3 2 6
1 6 5 4 3 2
2 1 3 4 5 6
2 1 3 4 6 5
2 1 3 5 4 6
2 1 3 6 5 4
2 1 4 3 5 6
2 1 4 3 6 5
2 1 5 4 3 6
2 1 6 5 4 3
3 2 1 4 5 6
3 2 1 4 6 5
3 2 1 5 4 6
3 2 1 6 5 4
4 3 2 1 5 6
4 3 2 1 6 5
5 4 3 2 1 6
6 5 4 3 2 1`,
	})
	testutil.AssertEqualStringCase(t, tc, 0, CF1508B)
}
