package main

func buildArray(a []int, _ int) (ans []string) {
	j := 1
	for _, v := range a {
		for ; j < v; j++ {
			ans = append(ans, "Push", "Pop")
		}
		ans = append(ans, "Push")
		j++
	}
	return
}
