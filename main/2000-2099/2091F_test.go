// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/2091/F
// https://codeforces.com/problemset/status/2091/problem/F?friends=on
func Test_cf2091F(t *testing.T) {
	testCases := [][2]string{
		{
			`3
3 4 1
XX#X
#XX#
#X#X
3 4 2
XX#X
#XX#
#X#X
3 1 3
X
X
#`,
			`2
60
0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2091F)
}
