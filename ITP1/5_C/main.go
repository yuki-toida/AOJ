package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Print chess
func main() {
	sc := bufio.NewScanner(os.Stdin)
	s := ""
	for {
		sc.Scan()
		ary := strings.Split(strings.TrimSpace(sc.Text()), " ")
		H, _ := strconv.Atoi(ary[0])
		W, _ := strconv.Atoi(ary[1])
		if H == 0 && W == 0 {
			break
		}
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				if i%2 == 1 {
					if j%2 == 1 {
						s += "#"
					} else {
						s += "."
					}
				} else {
					if j%2 == 1 {
						s += "."
					} else {
						s += "#"
					}
				}
			}
			s += "\n"
		}
		s += "\n"
	}
	fmt.Print(s)
}
