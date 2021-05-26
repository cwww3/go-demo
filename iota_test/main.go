package main

import "fmt"

const a = iota //0

const (
	b = iota // 0
)

const (
	c = 0    //0
	d = iota //1
	e        //2
	f = "hello"
	// nothing
	g           //4
	h    = iota //5
	i           //6
	j    = 0
	k                 // k =0  0
	l, m = iota, iota //9,9
	n, o              // n,o  = iota, iota  10,10

	p = iota + 1
	q
	_
	r = iota * iota //14*14
	s               // s =iota * iota 15*15
	t = r
	u //u=r
	v = 1 << iota
	w
	x         = iota * 0.01
	y float32 = iota * 0.01
	z
)

func main() {
	fmt.Printf("a : %T = %v\n", a, a)
	fmt.Printf("b : %T = %v\n", b, b)
	fmt.Printf("c : %T = %v\n", c, c)
	fmt.Printf("d : %T = %v\n", d, d)
	fmt.Printf("e : %T = %v\n", e, e)
	fmt.Printf("f : %T = %v\n", f, f)
	fmt.Printf("g : %T = %v\n", g, g)
	fmt.Printf("h : %T = %v\n", h, h)
	fmt.Printf("i : %T = %v\n", i, i)
	fmt.Printf("j : %T = %v\n", j, j)
	fmt.Printf("k : %T = %v\n", k, k)
	fmt.Printf("l : %T = %v\n", l, l)
	fmt.Printf("m : %T = %v\n", m, m)
	fmt.Printf("n : %T = %v\n", n, n)
	fmt.Printf("o : %T = %v\n", o, o)
	fmt.Printf("p : %T = %v\n", p, p)
	fmt.Printf("q : %T = %v\n", q, q)
	fmt.Printf("r : %T = %v\n", r, r)
	fmt.Printf("s : %T = %v\n", s, s)
	fmt.Printf("t : %T = %v\n", t, t)
	fmt.Printf("u : %T = %v\n", u, u)
	fmt.Printf("v : %T = %v\n", v, v)
	fmt.Printf("w : %T = %v\n", w, w)
	fmt.Printf("x : %T = %v\n", x, x)
	fmt.Printf("y : %T = %v\n", y, y)
	fmt.Printf("z : %T = %v\n", z, z)
}
