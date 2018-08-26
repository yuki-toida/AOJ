package main

import (
	"fmt"
)

type dice struct {
	t, s, e, w, n, b int
}

// サイコロ1
func main() {
	var t, s, e, w, n, b int
	fmt.Scan(&t, &s, &e, &w, &n, &b)
	d := dice{t: t, s: s, e: e, w: w, n: n, b: b}

	var cmds string
	fmt.Scan(&cmds)
	for _, v := range cmds {
		switch string(v) {
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
	fmt.Println(d.t)
}
