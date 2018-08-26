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

// サイコロ2
func main() {
	var t, s, e, w, n, b int
	fmt.Scan(&t, &s, &e, &w, &n, &b)
	d := dice{t: t, s: s, e: e, w: w, n: n, b: b}

	var q int
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var qt, qs int
		fmt.Scan(&qt, &qs)
		for _, v := range "EEENEEENEEESEEESEEENEEEN" {
			if d.t == qt && d.s == qs {
				fmt.Println(d.e)
				break
			}
			dir := string(v)
			d.rotate(dir)
		}
	}
}
