package main

func canReach(arr []int, start int) bool {
	n := len(arr)
    vis := make([]bool,n )
    var f func(i int) bool
    f = func(i int) bool {
    	if i<0 || i>=n {
			return false
		}
		if arr[i] == 0 {
			return true
		}
		if vis[i]{
			return false
		}
		vis[i]=true
		return f(i+arr[i]) || f(i-arr[i])
	}
	return f(start)
}
