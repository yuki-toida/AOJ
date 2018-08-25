package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Circle in rectangle
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ary := strings.Split(strings.TrimSpace(sc.Text()), " ")
	W, _ := strconv.Atoi(ary[0])
	H, _ := strconv.Atoi(ary[1])
	x, _ := strconv.Atoi(ary[2])
	y, _ := strconv.Atoi(ary[3])
	r, _ := strconv.Atoi(ary[4])
	if r <= x && r <= y && x+r <= W && y+r <= H {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
