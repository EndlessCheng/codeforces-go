package main

// github.com/EndlessCheng/codeforces-go
type Robot struct{}

var w, h, step int

func Constructor(width, height int) (_ Robot) {
	w, h, step = width, height, -1
	return
}

func (Robot) Move(num int) {
	if step < 0 {
		step = 0
	}
	step = (step + num) % ((w + h - 2) * 2)
}

func get() (x, y int, dir string) {
	switch {
	case step < 0:
		return 0, 0, "East"
	case step == 0:
		return 0, 0, "South"
	case step < w:
		return step, 0, "East"
	case step < w+h-1:
		return w - 1, step - w + 1, "North"
	case step < w*2+h-2:
		return w*2 + h - 3 - step, h - 1, "West"
	default:
		return 0, (w+h-2)*2 - step, "South"
	}
}

func (Robot) GetPos() []int  { x, y, _ := get(); return []int{x, y} }
func (Robot) GetDir() string { _, _, d := get(); return d }
