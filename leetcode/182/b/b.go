package main

func numTeams(a []int) (ans int) {
	n := len(a)
	for i, ai := range a {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if ai < a[j] && a[j] < a[k] || ai > a[j] && a[j] > a[k] {
					ans++
				}
			}
		}
	}
	return
}
