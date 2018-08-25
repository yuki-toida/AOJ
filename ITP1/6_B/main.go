package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 不足しているカードの発見
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	m := map[string][]bool{
		"S": []bool{false, false, false, false, false, false, false, false, false, false, false, false, false},
		"H": []bool{false, false, false, false, false, false, false, false, false, false, false, false, false},
		"C": []bool{false, false, false, false, false, false, false, false, false, false, false, false, false},
		"D": []bool{false, false, false, false, false, false, false, false, false, false, false, false, false},
	}

	for i := 0; i < n; i++ {
		sc.Scan()
		ary := strings.Split(sc.Text(), " ")
		s := ary[0]
		v, _ := strconv.Atoi(ary[1])
		mv, _ := m[s]
		mv[v-1] = true
	}

	s, _ := m["S"]
	h, _ := m["H"]
	c, _ := m["C"]
	d, _ := m["D"]
	str := print("S", s)
	str += print("H", h)
	str += print("C", c)
	str += print("D", d)
	fmt.Print(str)
}

func print(suit string, ary []bool) string {
	s := ""
	for i, v := range ary {
		if !v {
			s += fmt.Sprintf("%v %v\n", suit, i+1)
		}
	}
	return s
}
