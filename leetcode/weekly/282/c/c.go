package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minimumTime(time []int, totalTrips int) int64 {
	min := time[0]
	for _, v := range time[1:] {
		if v < min {
			min = v
		}
	}
	return int64(sort.Search(totalTrips*min, func(maxTime int) bool {
		tot := 0
		for _, t := range time {
			tot += maxTime / t
			if tot >= totalTrips {
				return true
			}
		}
		return false
	}))
}
