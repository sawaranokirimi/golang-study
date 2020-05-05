package main

import "fmt"

func main() {
	const MAX int = 1000

	// 1000 は含まないので、1000 - 1 してから計算
	answer := Sum(Filter(List(MAX - 1)))

	fmt.Printf("Ans. %v\n", answer)
	fmt.Println("done")
}

// 配列の各要素の値を足し合わせる
func Sum(xs []int) int {
	if len(xs) == 1 {
		return xs[0]
	}
	return xs[0] + Sum(xs[1:])
}

// 3 or 5 で割り切れない要素を配列から除外
func Filter(xs []int) []int {
	if len(xs) == 0 {
		return []int{}
	}

	if xs[0]%3 == 0 || xs[0]%5 == 0 {
		return append(Filter(xs[1:]), xs[0])
	}

	return Filter(xs[1:])
}

// 10 なら [1, 2, 3, ... 10] を返す
func List(n int) []int {
	if n == 1 {
		return []int{1}
	}
	return append(List(n-1), n)
}
