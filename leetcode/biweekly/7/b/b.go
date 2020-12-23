package main

import "strings"

// github.com/EndlessCheng/codeforces-go
type FileSystem struct{}

type node struct {
	ch  map[string]*node
	val int
}

var rt *node

func Constructor() (_ FileSystem) {
	rt = &node{ch: map[string]*node{}}
	return
}

func (FileSystem) CreatePath(path string, value int) (ans bool) {
	sp := strings.Split(path[1:], "/")
	n := len(sp)
	o := rt
	for _, s := range sp[:n-1] {
		if o.ch[s] == nil {
			return
		}
		o = o.ch[s]
	}
	end := sp[n-1]
	if o.ch[end] != nil {
		return
	}
	o.ch[end] = &node{map[string]*node{}, value}
	return true
}

func (FileSystem) Get(path string) (ans int) {
	o := rt
	for _, s := range strings.Split(path[1:], "/") {
		if o.ch[s] == nil {
			return -1
		}
		o = o.ch[s]
	}
	return o.val
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.CreatePath(path,value);
 * param_2 := obj.Get(path);
 */
