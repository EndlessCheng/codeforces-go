package main

import "slices"

var dis [26][26]int

func init() {
	const column = 6
	for i := range 26 {
		for j := range 26 {
			dis[i][j] = abs(i/column-j/column) + abs(i%column-j%column)
		}
	}
}

func minimumDistance(word string) int {
	f := [26]int{}
	for i := range len(word) - 1 {
		x, y := word[i]-'A', word[i+1]-'A'
		fy := f[y]
		for anotherFinger := range 26 {
			f[anotherFinger] = min(f[anotherFinger]+dis[x][y], fy+dis[x][anotherFinger])
		}
	}
	return slices.Min(f[:])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//

func minimumDistance1(word string) int {
	n := len(word)
	f := make([][26]int, n)

	for i := range n - 1 {
		x, y := word[i]-'A', word[i+1]-'A'
		for anotherFinger := range 26 {
			f[i+1][anotherFinger] = min(f[i][anotherFinger]+dis[y][x], f[i][y]+dis[anotherFinger][x])
		}
	}

	return slices.Min(f[n-1][:])
}
