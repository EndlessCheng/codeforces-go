package main

import "math"

type vecF struct{ x, y float64 }

func (a vecF) sub(b vecF) vecF     { return vecF{a.x - b.x, a.y - b.y} }
func (a vecF) len2() float64       { return a.x*a.x + a.y*a.y }
func (a vecF) dis2(b vecF) float64 { return a.sub(b).len2() }

func getCircleCenter(a, b vecF, r float64) vecF {
	mx, my := (a.x+b.x)/2, (a.y+b.y)/2
	d := math.Sqrt(r*r - a.dis2(b)/4)
	angle := math.Atan2(b.y-a.y, b.x-a.x)
	return vecF{mx + d*math.Sin(angle), my - d*math.Cos(angle)}
}

func numPoints(points [][]int, R int) (ans int) {
	ps := []vecF{}
	for _, p := range points {
		ps = append(ps, vecF{float64(p[0]), float64(p[1])})
	}
	r := float64(R)
	ans = 1
	const eps = 1e-8
	for i, a := range ps {
		for j, b := range ps {
			if j == i {
				continue
			}
			l := a.dis2(b)
			if l > 4*r*r+eps {
				continue
			}
			o := getCircleCenter(a, b, r)
			cnt := 0
			for _, p := range ps {
				if o.dis2(p) < r*r+eps {
					cnt++
				}
			}
			if cnt > ans {
				ans = cnt
			}
		}
	}
	return
}
