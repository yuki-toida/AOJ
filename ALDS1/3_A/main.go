package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	top int
	ary []int
}

func (s *stack) push(v int) {
	s.ary[s.top] = v
	s.top++
}

func (s *stack) pop() int {
	v := s.ary[s.top-1]
	s.top--
	return v
}

// スタック
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	s := stack{top: 0, ary: make([]int, len(inputs))}
	for _, v := range inputs {
		n, err := strconv.Atoi(v)
		if err == nil {
			s.push(n)
			continue
		}
		switch v {
		case "+":
			v1 := s.pop()
			v2 := s.pop()
			s.push(v2 + v1)
		case "-":
			v1 := s.pop()
			v2 := s.pop()
			s.push(v2 - v1)
		case "*":
			v1 := s.pop()
			v2 := s.pop()
			s.push(v2 * v1)
		}
	}
	fmt.Println(s.ary[0])
}
