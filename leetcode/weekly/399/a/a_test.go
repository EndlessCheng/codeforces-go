// Generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_a(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, numberOfPairs, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-399/problems/find-the-number-of-good-pairs-i/
// https://leetcode.cn/problems/find-the-number-of-good-pairs-i/