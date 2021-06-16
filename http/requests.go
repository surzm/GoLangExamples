package main

import (
	"fmt"
	"github.com/levigross/grequests"
	"log"
)

type resStruct struct {
	Args    struct{}          `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
}

func main() {
	resp, err := grequests.Get("http://httpbin.org/get", nil)
	// You can modify the request by passing an optional RequestOptions struct

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	fmt.Println(resp.String())

	//r := new(resStruct) // = &resStruct{}
	rr := resStruct{}
	if err := resp.JSON(&rr); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v", rr)
}
