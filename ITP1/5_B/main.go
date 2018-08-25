package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Print frame
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
				if i == 0 || i == H-1 {
					s += "#"
				} else {
					if j == 0 || j == W-1 {
						s += "#"
					} else {
						s += "."
					}
				}
			}
			s += "\n"
		}
		s += "\n"
	}
	fmt.Print(s)
}
