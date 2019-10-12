package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol1037D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for m := n - 1; m > 0; m-- {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	expectedOrder := make([]int, 0, n)
	visited := make([]bool, n+1)
	visited[1] = true
	for queue := []int{1}; len(queue) > 0; {
		v, queue = queue[0], queue[1:]
		expectedOrder = append(expectedOrder, v)
		sort.Ints(g[v])
		for _, w := range g[v] {
			if !visited[w] {
				visited[w] = true
				queue = append(queue, w)
			}
		}
	}

	//Fscan(in, &v)
	//if v != 1 {
	//	Fprint(out, "No")
	//	return
	//}

	actualOrder := make([]int, n)
	for i := range actualOrder {
		Fscan(in, &actualOrder[i])
	}
	l, r := 0, 1
	visited = make([]bool, n+1)
	visited[1] = true
	for queue := []int{1}; len(queue) > 0; {
		v, queue = queue[0], queue[1:]
		for _, w := range g[v] {
			if !visited[w] {
				visited[w] = true
				queue = append(queue, w)
			}
		}
	}

	for i, v := range expectedOrder {
		if v != actualOrder[i] {
			Fprint(out, "No")
			return
		}
	}
	Fprint(out, "Yes")
}

func main() {
	Sol1037D(os.Stdin, os.Stdout)
}
