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

// サイコロ4
func main() {
	var n int
	fmt.Scan(&n)
	dices := make([]dice, n)
	for i := 0; i < n; i++ {
		var t, s, e, w, n, b int
		fmt.Scan(&t, &s, &e, &w, &n, &b)
		dices[i] = dice{t: t, s: s, e: e, w: w, n: n, b: b}
	}

	same := false
	base := dices[0]
	for i := 1; i < n; i++ {
		if same {
			break
		}
		for _, v := range "EEENEEENEEESEEESEEENEEEN" {
			if base.equal(dices[i]) {
				same = true
				break
			}
			dices[i].rotate(string(v))
		}
	}

	if same {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
