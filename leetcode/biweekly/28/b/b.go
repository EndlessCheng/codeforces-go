package main

// github.com/EndlessCheng/codeforces-go
type SubrectangleQueries struct{}

var a [][]int

func Constructor(rectangle [][]int) (_ SubrectangleQueries) {
	a = rectangle
	return
}

func (SubrectangleQueries) UpdateSubrectangle(row1, col1, row2, col2, newValue int) {
	for _, r := range a[row1 : row2+1] {
		for j := col1; j <= col2; j++ {
			r[j] = newValue
		}
	}
}

func (SubrectangleQueries) GetValue(row, col int) int {
	return a[row][col]
}
