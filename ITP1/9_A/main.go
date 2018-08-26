package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 単語の検索
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	W := sc.Text()
	cnt := 0
	for {
		sc.Scan()
		s := sc.Text()
		if s == "END_OF_TEXT" {
			break
		}
		if W == strings.ToLower(s) {
			cnt++
		}
	}
	fmt.Println(cnt)
}
