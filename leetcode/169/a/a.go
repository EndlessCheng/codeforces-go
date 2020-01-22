package main

func sumZero(n int) (ans []int) {
    if n &1 ==0 {
    	m:= n/2
    	for i :=1;i<=m;i++ {
    		ans = append(ans, i,-i)
		}
	} else {
		ans = append(ans, 0)
		m:=n/2
		for i :=1;i<=m;i++ {
			ans = append(ans, i,-i)
		}
	}
	return
}
