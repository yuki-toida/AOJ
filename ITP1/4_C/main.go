package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Simple calculater
func main() {
	sc := bufio.NewScanner(os.Stdin)
	s := ""
	for {
		sc.Scan()
		ary := strings.Split(strings.TrimSpace(sc.Text()), " ")
		op := ary[1]
		if op == "?" {
			break
		}
		a, _ := strconv.Atoi(ary[0])
		b, _ := strconv.Atoi(ary[2])
		switch op {
		case "+":
			s += fmt.Sprintf("%v\n", a+b)
		case "-":
			s += fmt.Sprintf("%v\n", a-b)
		case "*":
			s += fmt.Sprintf("%v\n", a*b)
		case "/":
			s += fmt.Sprintf("%v\n", a/b)
		}
	}
	fmt.Print(s)
}
