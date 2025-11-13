package main

func buildArray(target []int, _ int) (ans []string) {
	mx := target[len(target)-1]
	i := 0
	for x := 1; x <= mx; x++ {
		ans = append(ans, "Push") // 先把 x 入栈
		if target[i] == x { // x 是我们要的数
			i++
		} else { // x 不是我们要的数，出栈
			ans = append(ans, "Pop")
		}
	}
	return
}
