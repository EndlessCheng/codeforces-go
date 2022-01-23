package main

// github.com/EndlessCheng/codeforces-go
func maximumGood(statements [][]int) (ans int) {
next:
	for i := 1; i < 1<<len(statements); i++ {
		cnt := 0 // i 中好人个数
		for j, row := range statements {
			if i>>j&1 == 1 { // 枚举 i 中的好人 j
				for k, st := range row { // 枚举 j 的所有陈述 st
					if st < 2 && st != i>>k&1 { // 该陈述与实际情况矛盾
						continue next
					}
				}
				cnt++
			}
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return
}
