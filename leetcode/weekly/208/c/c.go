package main

// github.com/EndlessCheng/codeforces-go
type ThroneInheritance struct{}

func Constructor(kingName string) (t ThroneInheritance) {
	root = kingName
	g = map[string][]string{root: {}}
	died = map[string]bool{}
	return
}

var root string
var g map[string][]string
var died map[string]bool

func (ThroneInheritance) Birth(parentName string, childName string) {
	g[parentName] = append(g[parentName], childName)
}

func (t *ThroneInheritance) Death(name string) {
	died[name] = true
}

func (ThroneInheritance) GetInheritanceOrder() (ans []string) {
	var f func(string)
	f = func(v string) {
		if !died[v] {
			ans = append(ans, v)
		}
		for _, w := range g[v] {
			f(w)
		}
	}
	f(root)
	return
}
