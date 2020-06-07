package main

type BrowserHistory struct {
}

var (
	s []string
	i int
)

func Constructor(home string) (b BrowserHistory) {
	s = []string{home}
	i = 0
	return
}

func (*BrowserHistory) Visit(url string) {
	s = append(s[:i+1], url)
	i = len(s) - 1
}

func (*BrowserHistory) Back(steps int) (ans string) {
	i -= steps
	if i < 0 {
		i = 0
	}
	return s[i]
}

func (*BrowserHistory) Forward(steps int) (ans string) {
	i += steps
	if i >= len(s) {
		i = len(s) - 1
	}
	return s[i]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
