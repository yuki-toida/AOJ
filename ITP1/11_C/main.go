package main

import (
	"fmt"
)

type dice struct {
	t, s, e, w, n, b int
}

func (d *dice) rotate(s string) {
	switch s {
	case "E":
		d.t, d.e, d.w, d.b = d.w, d.t, d.b, d.e
	case "W":
		d.t, d.e, d.w, d.b = d.e, d.b, d.t, d.w
	case "S":
		d.t, d.s, d.b, d.n = d.n, d.t, d.s, d.b
	case "N":
		d.t, d.s, d.b, d.n = d.s, d.b, d.n, d.t
	}
}

func (d *dice) equal(t dice) bool {
	return d.t == t.t && d.s == t.s && d.e == t.e && d.w == t.w && d.n == t.n && d.b == t.b
}

// サイコロ3
func main() {
	var t1, s1, e1, w1, n1, b1 int
	fmt.Scan(&t1, &s1, &e1, &w1, &n1, &b1)
	d1 := dice{t: t1, s: s1, e: e1, w: w1, n: n1, b: b1}

	var t2, s2, e2, w2, n2, b2 int
	fmt.Scan(&t2, &s2, &e2, &w2, &n2, &b2)
	d2 := dice{t: t2, s: s2, e: e2, w: w2, n: n2, b: b2}

	s := "No"
	for _, v := range "EEENEEENEEESEEESEEENEEEN" {
		if d1.equal(d2) {
			s = "Yes"
			break
		}
		d1.rotate(string(v))
	}
	fmt.Println(s)
}
