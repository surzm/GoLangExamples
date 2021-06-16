package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {

	// using standard library "flag" package
	flag.Int("f", 1234, "help message for flagname")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	i := viper.GetInt("f") // retrieve value from viper

	fmt.Println(i)
	fmt.Println(pflag.Args())
}
