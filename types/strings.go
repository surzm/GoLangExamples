package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	structEx := struct {
		Weight int
		Height int
	}{}

	fmt.Println("Some string")
	fmt.Printf("Some string %s %d %+v", "string", 88, structEx)

	str1 := fmt.Sprintf("世界Some string %s %d %+v", "string", 88, structEx)
	fmt.Println(str1)

	fmt.Println(str1[:2])
	unirune := str1[:3] // 世 - 3 байта

	fmt.Println(unirune, len(unirune), utf8.RuneCountInString(unirune))
	fmt.Println(str1, len(str1), utf8.RuneCountInString(str1))

	//for i, v := range str1 {
	//	fmt.Println(v)
	//	fmt.Printf("%d %s \n", i, string(v))
	//}
	fmt.Printf("%s %s \n", "Helloo", "world")
	fmt.Println(unirune + " hello")
	//strings

}
