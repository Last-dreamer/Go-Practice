package main

import (
	"fmt"
	"math/rand"
	"time"
)


func roll(side int) int {
	return rand.Intn(side)+ 1
}
func main() {
	rand.Seed(time.Now().UnixNano())	

	dice, sides := 2,6
	roles := 2
	for r:=1;r<=roles;r++ {
		sum:= 0
		for d:=1;d<=dice;d++ {
		 rolled := roll(sides)
		 sum = rolled
		 fmt.Println("Roll # ",r , "die #", d, ":", rolled)
		}

		fmt.Println("Total Rolled : ", sum)

		switch sum:=sum; {
		case sum ==  2 && dice == 2:
			fmt.Println("Snack Eye")
		case sum == 7:
			fmt.Println("Lucky seven ")
		case sum % 2 ==0:
			fmt.Println("Even") 
		case sum % 2 ==1:
			fmt.Println("Odd") 

		}
	}

}