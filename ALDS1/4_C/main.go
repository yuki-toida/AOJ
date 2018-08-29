package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 辞書
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	m := make(map[string]bool, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		ary := strings.Split(sc.Text(), " ")
		key := ary[1]
		switch ary[0] {
		case "insert":
			if _, ok := m[key]; !ok {
				m[key] = true
			}
		case "find":
			if _, ok := m[key]; ok {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		}
	}
}
