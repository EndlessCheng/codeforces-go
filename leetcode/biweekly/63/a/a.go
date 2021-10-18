package main

import "sort"

/* 排序+贪心

由于座位和学生数相同，一个萝卜一个坑，将座位和学生位置排序后，第 $i$ 个学生可以对应第 $i$ 个座位。

由于交换任意两个学生对应的座位不会产生更少的移动次数（可以画一画，证明略），所以上述对应关系可以产生最少移动次数，累加位置之差即为答案。

*/

// github.com/EndlessCheng/codeforces-go
func minMovesToSeat(seats, students []int) (ans int) {
	sort.Ints(seats)
	sort.Ints(students)
	for i, p := range seats {
		ans += abs(p - students[i])
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
