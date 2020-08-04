package _00_199

import (
	"bufio"
	. "fmt"
	"io"
)

type pair190C struct {
	isInt bool
	a, b  *pair190C
}

func (p *pair190C) read(in *bufio.Reader) bool {
	var s string
	if n, _ := Fscan(in, &s); n == 0 {
		return false
	}
	if s == "int" {
		p.isInt = true
		return true
	}
	p.a = &pair190C{}
	if !p.a.read(in) {
		return false
	}
	p.b = &pair190C{}
	if !p.b.read(in) {
		return false
	}
	return true
}

func (p *pair190C) print(out *bufio.Writer) {
	if p.isInt {
		Fprint(out, "int")
		return
	}
	Fprint(out, "pair<")
	p.a.print(out)
	Fprint(out, ",")
	p.b.print(out)
	Fprint(out, ">")
}

func Sol190C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	const e = "Error occurred"
	var n int
	Fscan(in, &n)
	p := &pair190C{}
	if !p.read(in) {
		Fprint(out, e)
		return
	}
	var s string
	if n, _ := Fscan(in, &s); n > 0 {
		Fprint(out, e)
		return
	}
	p.print(out)
}

//func main() {
//	Sol190C(os.Stdin, os.Stdout)
//}
