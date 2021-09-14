package main

import (
	"sort"
	"strconv"
)

func displayTable(orders [][]string) (ans [][]string) {
	uniqueInPlace := func(a []string) []string {
		n := len(a)
		j := 0
		for i := 1; i < n; i++ {
			if a[j] != a[i] {
				j++
				a[j] = a[i]
			}
		}
		return a[:j+1]
	}

	foods := []string{}
	tables := [501]map[string]int{}
	for i := range tables {
		tables[i] = map[string]int{}
	}
	for _, order := range orders {
		tt, fd := order[1], order[2]
		t, _ := strconv.Atoi(tt)
		foods = append(foods, fd)
		tables[t][fd]++
	}

	sort.Strings(foods)
	foods = uniqueInPlace(foods)

	ans = append(ans, append([]string{"Table"}, foods...))
	for i, m := range tables {
		if len(m) > 0 {
			a := []string{strconv.Itoa(i)}
			for _, fd := range foods {
				a = append(a, strconv.Itoa(m[fd]))
			}
			ans = append(ans, a)
		}
	}
	return
}
