package main

import (
	"fmt"

	"play/array"
)

func main() {
	src := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(src)

	// 配列から配列を取り除く
	// list1 := []int{1, 2, 3, 4, 5, 6}
	// list2 := []int{1, 3, 5}
	// fmt.Println(array.Take(list1, list2))

	// 配列を逆順に並び替える
	// fmt.Println(src)
	// dst := array.Reverse(src)
	// fmt.Println(dst)
	// fmt.Println(src)

	// 配列を指定したサイズで分割する
	// dst := array.Split(src, 2)
	// fmt.Println(dst)
	// fmt.Println(src)

	// 配列を条件でフィルタリングする
	condisions := func(n int) bool {
		return n%2 == 0
	}
	dst := array.Filter(src, condisions)
	fmt.Println(dst)
	fmt.Println(src)
}
