package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	sampleIns := [][]string{{`"lee(t(c)o)de)"`}, {`"a)b(c)d"`}, {`"))(("`}, {`"(a(b(c)d)"`}}
	sampleOuts := [][]string{{`"lee(t(c)o)de"`}, {`"ab(c)d"`}, {`""`}, {`"a(b(c)d)"`}}
	if err := testutil.RunLeetCodeFunc(t, minRemoveToMakeValid, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
