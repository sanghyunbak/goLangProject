package main

import (
	"fmt"
	"github.com/sanghyunbak/goLangProject/config"
	"rsc.io/quote"
)

func main() {

	fmt.Println("hello world")

	config.Test()
	abc := config.ABC{1}
	abc.Print()
	quote.Hello()
}
