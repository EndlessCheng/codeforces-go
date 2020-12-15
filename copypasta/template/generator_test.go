package main

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

// 赛前看情况修改下 main.go  main_test.go

// https://codeforces.com/contest//problems
func TestGenCodeforcesContestTemplates(t *testing.T) {
	const contestID = ""
	const overwrite = false
	for {
		if err := GenCodeforcesContestTemplates(contestID, overwrite); err != nil {
			//t.Log(err)
		} else {
			break
		}
		time.Sleep(5 * time.Second)
	}
}

// https://codeforces.com/problemset/problem/1293/C
// https://codeforces.com/problemset/status/1291/problem/D
// https://codeforces.com/gym/102253/problem/C
// https://codeforces.com/gym/102253/status/C
func TestGenCodeforcesProblemTemplates(t *testing.T) {
	raw, err := ioutil.ReadFile("data.txt")
	if err != nil {
		t.Fatal(err)
	}
	problemURL := strings.TrimSpace(string(raw))
	if err := GenCodeforcesProblemTemplates(problemURL, true); err != nil {
		t.Fatal(err)
	}
}

// "../../misc/kickstart/2020//"
// "../../misc/luogu/contest//"
// "../../misc/nowcoder//"
func TestGenTemplates(t *testing.T) {
	const (
		problemNum = 4
		rootPath   = "../../misc/luogu/contest//"
		overwrite  = false
	)
	if err := GenTemplates(problemNum, rootPath, overwrite); err != nil {
		t.Fatal(err)
	}
}
