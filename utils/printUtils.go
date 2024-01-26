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

func PrintIntArray(array []int) {
	fmt.Print("[")
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i])
		fmt.Print(",\t")
	}
	fmt.Println("]")
}
