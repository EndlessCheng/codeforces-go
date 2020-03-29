package main

type UndergroundSystem struct {
	tin     map[int]map[string]int
	records map[string]map[string][2]int
}

func Constructor() (u UndergroundSystem) {
	u.tin = map[int]map[string]int{}
	u.records = map[string]map[string][2]int{}
	return
}

func (u *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	in, ok := u.tin[id]
	if !ok {
		in = map[string]int{}
	}
	in[stationName] = t
	u.tin[id] = in
}

func (u *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	for st, tin := range u.tin[id] {
		out, ok := u.records[st]
		if !ok {
			out = map[string][2]int{}
		}
		r := out[stationName]
		r[0] += t - tin
		r[1]++
		out[stationName] = r
		u.records[st] = out
	}
}

func (u *UndergroundSystem) GetAverageTime(startStation string, endStation string) (ans float64) {
	r := u.records[startStation][endStation]
	return float64(r[0]) / float64(r[1])
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
