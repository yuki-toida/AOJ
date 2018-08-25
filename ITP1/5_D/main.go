package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Structured programing
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	s := ""
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			s += fmt.Sprintf("%v ", i)
		} else {
			cnt := strings.Count(strconv.Itoa(i), "3")
			if 0 < cnt {
				s += fmt.Sprintf("%v ", i)
			}
		}
	}
	fmt.Println(" " + strings.TrimSpace(s))
}
