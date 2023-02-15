package main

import (
	"os"

	"github.com/cbot918/liby/ready"
	"github.com/cbot918/yo/src/yo"
)

// # yo
// src
// https://juejin.cn/post/6844903672082595847

func main(){
	r := ready.New()
	
	program := r.Get("./code.js")
	// fmt.Println(program)
	// program := " a234 world"

	args := os.Args

	// fmt.Println([]byte(" "))
	if len(args) == 1 {
		y := yo.New(program)
		// y.ListString()
		y.Run()
	}
	if len(args) == 2 {
		y := yo.New(args[1])
		y.Run()
	}

}