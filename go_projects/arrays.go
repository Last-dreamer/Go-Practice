package main

import "fmt"


type Products struct {
	price int 
	name string
}

func printstates(list [4]Products){
	var cost , totalItem int

	for i:=0;i<len(list);i++ {
		item := list[i]
		cost += item.price

		if item.name != "" {
			totalItem += 1
		}

	}

	fmt.Println("last item on the list: ", list[totalItem-1])
	fmt.Println("Total item :", totalItem)
	fmt.Println("Total Cost ", cost)
}

func main(){

	shoppingList := [4]Products{
		{1, "bannana"},
		{6, "Meat"},
		{2,"Salad"},
	}

	printstates(shoppingList)

	shoppingList[3] = Products{4, "Bread"}

	printstates(shoppingList)


}