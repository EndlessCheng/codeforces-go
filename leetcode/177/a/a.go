package main

import (
	"math"
	"time"
)

func daysBetweenDates(date1 string, date2 string) (ans int) {
	t1, _ := time.Parse("2006-01-02", date1)
	t2, _ := time.Parse("2006-01-02", date2)
	return int(math.Abs(t1.Sub(t2).Hours())) / 24
}
