package main

import (
	"fmt"
	"toolkits/slice"
)

func main() {
	fmt.Println(slice.Contains([]int{1, 2}, 1))
	fmt.Println(slice.Contains([]string{"1", "2"}, "2"))
	fmt.Println(slice.Contains([]string{"1", "2"}, "0"))
}
