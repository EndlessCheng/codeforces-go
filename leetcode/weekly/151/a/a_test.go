package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	sampleIns := [][]string{{`["alice,20,800,mtv","alice,50,100,beijing"]`}, {`["alice,20,800,mtv","alice,50,1200,mtv"]`}, {`["alice,20,800,mtv","bob,50,1200,mtv"]`}}
	sampleOuts := [][]string{{`["alice,20,800,mtv","alice,50,100,beijing"]`}, {`["alice,50,1200,mtv"]`}, {`["bob,50,1200,mtv"]`}}
	if err := testutil.RunLeetCodeFunc(t, invalidTransactions, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
