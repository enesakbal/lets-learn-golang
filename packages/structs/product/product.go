package product

type Product struct {
	ID         int
	Name       string
	Price      ProductPrice
	MerchantID int
}

type ProductPrice struct {
	Price    float64
	Currency string
	Discount float64
}
