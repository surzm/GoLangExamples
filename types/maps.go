package main

import "fmt"

func main() {
	mapEx := map[int]string{1: "1", 2: "2", 3: "3"}
	fmt.Println(mapEx[4])

	val, ok := mapEx[4]
	fmt.Println(val, ok) // извлечение

	//delete(mapEx,1) // удаление
	mapEx[4] = "4" // вставка

	//fmt.Println(&mapEx[1])

	for k, v := range mapEx {
		fmt.Println(k, v) // псевдорандом
	}

}
