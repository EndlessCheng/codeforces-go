package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	sampleIns := [][]string{{`["mobile","mouse","moneypot","monitor","mousepad"]`, `"mouse"`}, {`["havana"]`, `"havana"`}, {`["bags","baggage","banner","box","cloths"]`, `"bags"`}, {`["havana"]`, `"tatiana"`}}
	sampleOuts := [][]string{{`[["mobile","moneypot","monitor"],["mobile","moneypot","monitor"],["mouse","mousepad"],["mouse","mousepad"],["mouse","mousepad"]]`}, {`[["havana"],["havana"],["havana"],["havana"],["havana"],["havana"]]`}, {`[["baggage","bags","banner"],["baggage","bags","banner"],["baggage","bags"],["bags"]]`}, {`[[],[],[],[],[],[],[]]`}}
	if err := testutil.RunLeetCodeFunc(t, suggestedProducts, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
