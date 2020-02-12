package main

import (
	"testing"
)

// 生成比赛模板（需要先创建目录）
func TestGenContestTemplates(t *testing.T) {
	const contestID = "1300"
	const overwrite = false
	if err := GenContestTemplates(contestID, overwrite); err != nil {
		t.Fatal(err)
	}
}

// 生成单道题目的模板（Codeforces）
// https://codeforces.com/problemset/problem/1293/C
// https://codeforces.com/problemset/status/1291/problem/D
// https://codeforces.com/gym/102253/problem/C
func TestGenCodeforcesNormalTemplates(t *testing.T) {
	const problemURL = "https://codeforces.com/problemset/problem/1293/C"
	if err := GenCodeforcesNormalTemplates(problemURL); err != nil {
		t.Fatal(err)
	}
}

// 生成单道题目的模板（非 Codeforces）
func TestGenNormalTemplates(t *testing.T) {
	const rootPath = "../../nowcoder/2720/"
	const overwrite = false
	if err := GenNormalTemplates(rootPath, overwrite); err != nil {
		t.Fatal(err)
	}
}
