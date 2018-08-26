package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// カードゲーム
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	tac := 0
	hac := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		ary := strings.Split(sc.Text(), " ")
		ta := ary[0]
		ha := ary[1]
		if ta < ha {
			hac += 3
		} else if ha < ta {
			tac += 3
		} else {
			tac++
			hac++
		}
	}
	fmt.Printf("%v %v\n", tac, hac)
}
