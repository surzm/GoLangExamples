package main

import "fmt"

type Person struct {
	Name     string `json:"name"`
	Old      int    `json:"old"`
	Surname  string `json:"surname"`
	Children []Child
	Man
}

type Child struct {
	Name string
	Old  int
}

type Man struct {
	BeardColor string
	//Cars []string
}

func (p *Person) AddChild(ch Child) {
	p.Children = append(p.Children, ch)
	p.Surname = "Ivanov"
}

func (p Person) ChildCount() int {
	return len(p.Children)
}

type PersInterface interface {
	//AddChild(child Child)
	ChildCount() int
}

func testFunc(p PersInterface) {

}
func main() {
	per := Person{Children: []Child{
		{
			Name: "Vova",
			Old:  12,
		},
		{
			Name: "Liza",
			Old:  3,
		},
	}}

	//per.addChild(Child{
	//	Name: "333",
	//	Old:  33,
	//})

	testFunc(per)
	fmt.Printf("%+v", per)
	fmt.Println(per.ChildCount())
}
