package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][3]string{
		{
			`["WordFilter","f"]`,
			`[[["apple"]],["a","e"]]`,
			`[null,0]`,
		},
		{
			`["WordFilter","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f","f"]`,
			`[[["pop"]],["",""],["","p"],["","op"],["","pop"],["p",""],["p","p"],["p","op"],["p","pop"],["po",""],["po","p"],["po","op"],["po","pop"],["pop",""],["pop","p"],["pop","op"],["pop","pop"],["",""],["","p"],["","gp"],["","pgp"],["p",""],["p","p"],["p","gp"],["p","pgp"],["pg",""],["pg","p"],["pg","gp"],["pg","pgp"],["pgp",""],["pgp","p"],["pgp","gp"],["pgp","pgp"]]`,
			`[null,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-1,-1,0,0,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1]`,
		},
		{
			`["WordFilter","f","f","f","f","f","f","f","f","f","f"]`,
			`[[["abbbababbb","baaabbabbb","abababbaaa","abbbbbbbba","bbbaabbbaa","ababbaabaa","baaaaabbbb","babbabbabb","ababaababb","bbabbababa"]],["","abaa"],["babbab",""],["ab","baaa"],["baaabba","b"],["abab","abbaabaa"],["","aa"],["","bba"],["","baaaaabbbb"],["ba","aabbbb"],["baaa","aabbabbb"]]`,
			`[null,5,7,2,1,5,5,3,6,6,1]`,
		},
	}
	if err := testutil.RunLeetCodeClassWithExamples(t, Constructor, examples, 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.com/contest/weekly-contest-by-app-academy/problems/prefix-and-suffix-search/
// 745
