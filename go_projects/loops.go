package main

import "fmt"

func loops() int {
	var a int
	for i:=0; i<10;i++ {
		a = i
     fmt.Println("looping through", i)
	}

 return a
}
 