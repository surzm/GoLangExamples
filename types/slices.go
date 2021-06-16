package main

import "fmt"

func main() {
	var sliceEx []string
	initSl := make([]string, 0) // возвращает инициализированный массив
	newSl := new([]string)      // возвращает указатель на неинициализированный массив

	fmt.Println(sliceEx == nil) // true
	fmt.Println(initSl == nil)  // false
	fmt.Println(*newSl == nil)  // true
	fmt.Println(initSl)
	fmt.Println(newSl)

	capSl := make([]int, 10)
	fmt.Println(capSl)
	capSl = make([]int, 0, 10)
	capSl = append(capSl, 1)

	fmt.Println(capSl[0])
	fmt.Println(capSl)

	sl1 := []int{0, 1, 2, 3, 4, 5}

	part1 := sl1[4:6]
	fmt.Println(part1) // [4,5]
	//part1 = append(part1, 6)
	part1[1] = 6
	fmt.Println(part1)
	fmt.Println(sl1)
}
