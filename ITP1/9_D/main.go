package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 文字列変換
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())
	for i := 0; i < q; i++ {
		sc.Scan()
		ary := strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(ary[1])
		b, _ := strconv.Atoi(ary[2])
		switch ary[0] {
		case "replace":
			str = str[:a] + ary[3] + str[b+1:]
		case "reverse":
			str = str[:a] + reverse(str[a:b+1]) + str[b+1:]
		case "print":
			fmt.Printf("%v\n", str[a:b+1])
		}
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
