package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type stack struct {
	ary []interface{}
}

func (s *stack) push(v interface{}) {
	s.ary = append(s.ary, v)
}

func (s *stack) pop() interface{} {
	i := len(s.ary) - 1
	v := s.ary[i]
	s.ary = s.ary[:i]
	return v
}

type area struct {
	index, value int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s1 := stack{}
	s2 := stack{}
	for i, s := range sc.Text() {
		switch string(s) {
		case "\\":
			s1.push(i)
		case "/":
			if len(s1.ary) <= 0 {
				continue
			}
			sindex := s1.pop().(int)
			total := 0
			for {
				if len(s2.ary) <= 0 {
					break
				}
				sarea := s2.pop().(area)
				if sarea.index <= sindex {
					s2.push(area{index: sarea.index, value: sarea.value})
					break
				}
				total += sarea.value
			}
			s2.push(area{index: sindex, value: i - sindex + total})
		case "_":
		}
	}

	sum := 0
	s := strconv.Itoa(len(s2.ary))
	for _, v := range s2.ary {
		area := v.(area)
		sum += area.value
		s += " " + strconv.Itoa(area.value)
	}
	fmt.Println(sum)
	fmt.Println(s)
}
