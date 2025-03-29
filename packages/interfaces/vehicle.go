package interfaces

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

// !!
type Audi struct{} // Not implemented Vehicle interface

func (a Audi) String() string {
	return "Audi"
}
