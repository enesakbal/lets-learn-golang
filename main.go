package main

import (
	"fmt"

	hello_world "github.com/enesakbal/lets-learn-golang/packages/hello_world"
	structs "github.com/enesakbal/lets-learn-golang/packages/structs"
)

func main() {
	hello_world.SayHello("World from GO!")

	merchants := structs.GetAllMerchants()
	products := structs.GetAllProducts()

	fmt.Println(merchants)
	fmt.Println(merchants[0])
	fmt.Println(merchants[22]) //*Throws panic => runtime error: index out of range [22] with length 2

	fmt.Println(products)
	fmt.Println(products[0])

}
