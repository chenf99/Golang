package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

var inputName = flag.StringP("name", "n", "", "Input Your Name.")
var inputAge = flag.IntP("age", "a", 27, "Input Your Age")
var inputGender = flag.StringP("gender", "s", "female", "Input Your Gender")
var inputTest = flag.BoolP("test", "t", false, "test")
var inputFlagvar int

func Init() {
	flag.IntVarP(&inputFlagvar, "flagname", "i", 1234, "Help")
}
func main() {
	Init()
	flag.Parse()
	// func Args() []string
	// Args returns the non-flag command-line arguments.
	// func NArg() int
	// NArg is the number of arguments remaining after flags have been processed.
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println(*inputName == "")
	fmt.Println("name=", *inputName)
	fmt.Println("age=", *inputAge)
	fmt.Println("test=", *inputTest)
	fmt.Println("gender=", *inputGender)
	fmt.Println("flagname=", inputFlagvar)
}