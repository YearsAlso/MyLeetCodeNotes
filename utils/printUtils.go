package utils

import (
	"container/list"
	"fmt"
)

func PrintList(l *list.List) {
	fmt.Print("[")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
		if e.Next() != nil {
			fmt.Print(",\t")
		}
	}
	fmt.Println("]")
}
