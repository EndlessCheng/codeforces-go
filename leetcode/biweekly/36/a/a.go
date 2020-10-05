package main

// github.com/EndlessCheng/codeforces-go
type ParkingSystem struct {
}

func Constructor(b, m, s int) (p ParkingSystem) {
	left = [4]int{0, b, m, s}
	return
}

var left [4]int

func (ParkingSystem) AddCar(t int) (ans bool) {
	if left[t] == 0 {
		return
	}
	left[t]--
	return true
}
