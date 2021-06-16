package main

import "fmt"

var (
	b int
	c string = "string" // автоматический вывод типов
)

func main() {
	var a int
	var e = "str"
	d := 3 // var d = 3

	fmt.Println(a, b, c, d, e)

	/* for */
	sliceInt := []int{1, 2, 3, 4, 5, 6}
	for i, v := range sliceInt { // i, v - локальная переменная
		if i == 0 {
			v = 0
			//sliceInt[i] = 0
		}
		fmt.Println(v)
	}
	fmt.Println(sliceInt)

Loop1:
	for i := 0; i < 5; i++ {
		switch i {
		case 1, 4:
			fmt.Println(1) // break автоматический
		case 2:
			fmt.Println(2)
		//fallthrough
		case 3:
			fmt.Println("break loop")
			break Loop1
		default:
			fmt.Println("some")
		}
	}

	// variadic func
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	sum(1, 2, 3, 4)
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
