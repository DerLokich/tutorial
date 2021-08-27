package main

import (
	"fmt"
)

func main() {
	var s1 []int
	fmt.Println("пустой слайс", s1)
	s1 = append(s1, 100) //добавление элемента в слайс
	fmt.Println("уже не пустой слайс", s1)
	fmt.Println("длина слайса", len(s1))

	fmt.Println("длина внутреннего массива в слайсе", s1, cap(s1))

	s1 = append(s1, 102)
	fmt.Println("длина внутреннего массива в слайсе", s1, cap(s1))
	s1 = append(s1, 103)
	s1 = append(s1, 104)
	fmt.Println("длина внутреннего массива в слайсе", s1, cap(s1))
	s1 = append(s1, 105)
	fmt.Println("длина внутреннего массива в слайсе", s1, cap(s1))
	s1 = append(s1, 106)
	fmt.Println("длина внутреннего массива в слайсе", s1, cap(s1))

	sl2 := []int{10, 20, 30}
	fmt.Println(sl2, len(sl2), cap(sl2))

	//добавить слайс в слайс
	s1 = append(s1, sl2...)
	fmt.Println(s1)

	var slm [][]int
	slm = append(slm, sl2)
	fmt.Println(slm)

	slice3 := make([]int, 10)
	slice3 = append(slice3, 1)
	fmt.Println(slice3, len(slice3), cap(slice3))
}
