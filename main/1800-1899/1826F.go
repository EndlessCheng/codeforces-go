package main

import (
	. "fmt"
	"sort"
)

// https://space.bilibili.com/206214
func run26() {
	var T, n int
	var drop string
	for Scan(&T); T > 0; T-- {
		Scan(&n)
		Println("? 0 1 0")
		x0 := make([]float64, n)
		for i := range x0 {
			Scan(&x0[i], &drop)
		}
		Println("? 1 0 0")
		y0 := make([]float64, n)
		for i := range y0 {
			Scan(&drop, &y0[i])
		}
		Println("? 0.25 -50.25 0")
		xs := make([]float64, n)
		for i := range xs {
			Scan(&xs[i], &drop)
		}
		sort.Float64s(xs)
		Print("!")
		for _, x0 := range x0 {
			for _, y0 := range y0 {
				x := (201*x0 + y0) * 201 / (201*201 + 1)
				i := sort.SearchFloat64s(xs, x)
				if i > 0 && x-xs[i-1] < 5e-4 || i < n && xs[i]-x < 5e-4 {
					Printf(" %.5f %.5f", x0, y0)
				}
			}
		}
		Println()
	}
}

//func main() { run26() }
