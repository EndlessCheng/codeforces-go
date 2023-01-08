package main

// https://space.bilibili.com/206214
type DataStream struct{ value, k, cnt int }

func Constructor(value, k int) DataStream {
	return DataStream{value, k, 0}
}

func (d *DataStream) Consec(num int) bool {
	if num == d.value {
		d.cnt++
	} else {
		d.cnt = 0
	}
	return d.cnt >= d.k
}
