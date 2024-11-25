package main

import (
	"fmt"
	"iterators/slice"
)

func main() {
	arr := []int{1, 2, 3, 456}
	res := slice.From(arr).
		Reverse().
		ForEach(func(i int, v int) {
			fmt.Printf("%v ", v)
			if i == len(arr)-1 {
				fmt.Println()
			}
		}).
		Fill(1, 0, len(arr)-1).
		ForEach(func(i int, v int) {
			fmt.Printf("%v ", v)
			if i == len(arr)-1 {
				fmt.Println()
			}
		}).
		Map(func(i int, v int) int {
			return v*i + 3
		}).
		ForEach(func(i int, v int) {
			fmt.Printf("%v ", v)
			if i == len(arr)-1 {
				fmt.Println()
			}
		}).
		Filter(func(v int) bool {
			return v%2 == 0
		}).
		Collect()

	fmt.Println("Result:", res)

	idx, found := slice.From(res).Find(func(v int) bool { return v%2 == 0 })
	fmt.Println("Result find:", idx, found)

	everyRes := slice.From(res).Every(func(v int) bool { return v == 0 })
	fmt.Println("Result every:", everyRes)
}
