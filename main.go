package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		panic("Недостаточно аргументов командной строки")
	}
	function := os.Args[1]
	var intVar []int
	for i := 2; i < len(os.Args); i++ {
		if s, err := strconv.Atoi(os.Args[i]); err == nil {
			intVar = append(intVar, s)
		}
	}
	if len(intVar) == 0 {
		panic("Нет целых чисел для сортировки")
	}
	switch function {
	case "bubble":
		fmt.Println(bubble(intVar))
	case "qsort":
		fmt.Println(qsort(intVar))
	case "selection":
		fmt.Println(selection(intVar))
	default:
		fmt.Println("К сожалению, такой сортировки в данной программе нет.")
	}
}

func bubble(arr []int) []int {
	var flag int
	for j := 1; j <= len(arr)-1; j++ {
		flag = 0
		for i := 0; i <= len(arr)-1-j; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
	}
	return arr
}

func selection(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		var min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

func more(arr []int, i int) []int {
	var newArr []int
	for _, val := range arr {
		if val > i {
			newArr = append(newArr, val)
		}
	}
	return newArr
}
func less(arr []int, i int) []int {
	var newArr []int
	for _, val := range arr {
		if val < i {
			newArr = append(newArr, val)
		}
	}
	return newArr
}
func qsort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	pivot := arr[0]
	lessArr := less(arr[1:], pivot)
	moreArr := more(arr[1:], pivot)

	return append(append(qsort(lessArr), pivot), qsort(moreArr)...)
}
