package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dlist struct {
	value int
	prev  *dlist
	next  *dlist
}

var head, last *dlist

// 連結リスト
func main() {
	var (
		str string
		ary []string
		x   int
	)

	head = new(dlist)
	last = new(dlist)
	head.next = last
	last.prev = head

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		str = sc.Text()
		switch str {
		case "deleteFirst":
			head.next = head.next.next
			head.next.prev = head
		case "deleteLast":
			last.prev.prev.next = last
			last.prev = last.prev.prev
		default:
			ary = strings.Split(str, " ")
			x, _ = strconv.Atoi(ary[1])
			switch ary[0] {
			case "insert":
				insert(x)
			case "delete":
				delete(x)
			}
		}
	}

	dl := head.next
	fmt.Print(dl.value)
	dl = dl.next
	for {
		if dl == last {
			break
		}
		fmt.Print(" ", dl.value)
		dl = dl.next
	}
	fmt.Println()
}

func insert(v int) {
	dl := new(dlist)
	dl.value = v
	dl.prev = head
	dl.next = head.next

	head.next = dl
	dl.next.prev = dl
}

func delete(v int) {
	dl := head.next
	for dl != last {
		if dl.value == v {
			dl.prev.next = dl.next
			dl.next.prev = dl.prev
			return
		}
		dl = dl.next
	}
}
