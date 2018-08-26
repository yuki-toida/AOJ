package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 大文字と小文字の入れ替え
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	sr := ""
	for _, x := range s {
		xs := string(x)
		u := strings.ToUpper(xs)
		l := strings.ToLower(xs)
		if u == xs {
			sr += l
		} else if l == xs {
			sr += u
		} else {
			sr += xs
		}
	}
	fmt.Println(sr)
}
