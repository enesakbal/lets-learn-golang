package main

import (
	"fmt"

	hello_world "github.com/enesakbal/lets-learn-golang/packages/hello_world"
	interfaces "github.com/enesakbal/lets-learn-golang/packages/interfaces"
	structs "github.com/enesakbal/lets-learn-golang/packages/structs"
)

func main() {
	hello_world.SayHello("World from GO!")

	//* Struct Package
	merchants := structs.GetAllMerchants()
	products := structs.GetAllProducts()

	fmt.Println(merchants)
	fmt.Println(merchants[0])
	// fmt.Println(merchants[22]) //*Throws panic => runtime error: index out of range [22] with length 2

	fmt.Println(products)
	fmt.Println(products[0])

	//* Interface Package
	myTesla := interfaces.Tesla{}
	myMercedes := interfaces.Mercedes{}

	vehicleList := []any{
		myTesla,
		myMercedes,
		interfaces.Audi{}, //* Audi is not a vehicle
	}

	for i := 0; i < len(vehicleList); i++ {
		currentVehicle := vehicleList[i]

		if vehicle, ok := any(currentVehicle).(interfaces.Vehicle); ok {
			vehicleBrand := vehicle.GetVehicleModel()
			vehiclePrice := vehicle.GetPrice()
			fmt.Printf("%v is a Vehicle. Price is $%.2F\n", vehicleBrand, vehiclePrice)
		} else {
			fmt.Printf("%v is NOT a vehicle\n", currentVehicle)
		}
	}

}
