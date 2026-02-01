package main

// https://space.bilibili.com/206214
type RideSharingSystem struct {
	riders         []int
	drivers        []int
	seenRiders     map[int]bool
	canceledRiders map[int]bool
}

func Constructor() RideSharingSystem {
	return RideSharingSystem{
		seenRiders:     map[int]bool{},
		canceledRiders: map[int]bool{},
	}
}

func (r *RideSharingSystem) AddRider(riderId int) {
	r.riders = append(r.riders, riderId)
	r.seenRiders[riderId] = true
}

func (r *RideSharingSystem) AddDriver(driverId int) {
	r.drivers = append(r.drivers, driverId)
}

func (r *RideSharingSystem) MatchDriverWithRider() []int {
	// 弹出队列中的已取消乘客
	for len(r.riders) > 0 && r.canceledRiders[r.riders[0]] {
		r.riders = r.riders[1:]
	}
	// 没有乘客或者司机
	if len(r.riders) == 0 || len(r.drivers) == 0 {
		return []int{-1, -1}
	}
	// 配对
	ans := []int{r.drivers[0], r.riders[0]}
	r.riders = r.riders[1:]
	r.drivers = r.drivers[1:]
	return ans
}

func (r *RideSharingSystem) CancelRider(riderId int) {
	// 对于不存在的乘客，不能标记为取消
	if r.seenRiders[riderId] {
		r.canceledRiders[riderId] = true
	}
}
