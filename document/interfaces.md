# Interfaces
An interface is a collection of method signatures that define behavior. Unlike other languages, Go interfaces are satisfied implicitly—no explicit declaration is needed.

## Basic Interface Example

```go
type Vehicle interface {
	GetPrice() float32
	GetVehicleModel() string
}

type Tesla struct{}
type Mercedes struct{}

func (t Tesla) GetPrice() float32 {
	return 65.000
}

func (t Tesla) GetVehicleModel() string {
	return "Tesla Model Y"
}

func (t Mercedes) GetPrice() float32 {
	return 75.500
}

func (t Mercedes) GetVehicleModel() string {
	return "Mercedes A180"
}

func main() {
	myTesla := Tesla{}
	myMercedes := Mercedes{}

	vehicleList := []Vehicle{
		myTesla,
		myMercedes,
	}

	for i := 0; i < len(vehicleList); i++ {
		currentVehicle := vehicleList[i]

		fmt.Printf("my %v is $%0.3F\n", currentVehicle.GetVehicleModel(), currentVehicle.GetPrice())
	}
}
```
## How can you check that Tesla is Vehicle


## Type Assertion
A type assertion provides access to an interface value's underlying concrete value.
If i holds a T, then t will be the underlying value and ok will be true.

If not, ok will be false and t will be the zero value of type T, and no panic occurs.

Note the similarity between this syntax and that of reading from a map.

```go
// !!
type Audi struct{} // Not implemented Vehicle interface

func (a Audi) String() string {
	return "Audi"
}

func main() {
	var t Tesla = Tesla{}
    var a Audi := Audi{}


	vehicleList := []any{
		v,
        a,
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
```