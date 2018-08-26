package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 文字のカウント
func main() {
	sc := bufio.NewScanner(os.Stdin)
	alpha := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	m := map[string]int{}
	for _, v := range alpha {
		m[v] = 0
	}
	for {
		sc.Scan()
		s := sc.Text()
		if s == "" {
			break
		}
		for _, v := range s {
			char := strings.ToLower(string(v))
			_, ok := m[char]
			if !ok {
				continue
			}
			m[char]++
		}
	}
	s := ""
	for _, v := range alpha {
		s += fmt.Sprintf("%v : %v\n", v, m[v])
	}
	fmt.Print(s)
}
