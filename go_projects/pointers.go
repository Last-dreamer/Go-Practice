package main

import "fmt"


const (
	Active = true
	Inactive = false
)

// type alias to override the default datatype 
type SecurityTag bool


type Item struct {
	name string
	tag SecurityTag
}

func activate(tag *SecurityTag){
	*tag = Active
}

func deactivate(tag *SecurityTag){
	*tag = Inactive
}


func checkOut(item []Item){
	fmt.Println("checking out all the items ......")
	for i := 0; i < len(item);i++ {
		deactivate(&item[i].tag)
	}
}




func main(){

	shirt := Item{"Shirt", Active}
	pants := Item{"Pants", Active}
	 purse := Item{"purse", Active}
	watch := Item{"Watch", Active}

	items := []Item{shirt, pants, purse, watch}

	fmt.Println("intial :", items)

    deactivate(&items[0].tag)
	fmt.Println("after deact :", items)


}