package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// リング
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	sc.Scan()
	p := sc.Text()
	sr := strings.Replace(s+s, p, "", -1)
	if len([]rune(s+s)) == len([]rune(sr)) {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
