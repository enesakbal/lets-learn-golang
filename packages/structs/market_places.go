package structs

import (
	merchant "github.com/enesakbal/lets-learn-golang/packages/structs/merchant"
	product "github.com/enesakbal/lets-learn-golang/packages/structs/product"
)

var products = []product.Product{
	{
		ID:   1,
		Name: "MacBook Pro",
		Price: product.ProductPrice{
			Price:    25.999,
			Currency: "TRY",
		},
		MerchantID: 0,
	},
	{
		ID:   2,
		Name: "iPhone 14 Pro",
		Price: product.ProductPrice{
			Price:    100.999,
			Currency: "TRY",
		},
		MerchantID: 0,
	},
	{
		ID:   3,
		Name: "MacBook Pro",
		Price: product.ProductPrice{
			Price:    25.999,
			Currency: "TRY",
		},
		MerchantID: 0,
	},
	{
		ID:   4,
		Name: "Samsung S22",
		Price: product.ProductPrice{
			Price:    100.999,
			Currency: "TRY",
		},
		MerchantID: 1,
	},
	{
		ID:   5,
		Name: "Samsung 4K Monitor",
		Price: product.ProductPrice{
			Price:    100.999,
			Currency: "TRY",
		},
		MerchantID: 1,
	},
}

var merchants = []merchant.Merchant{
	{
		ID:       0,
		Name:     "Apple",
		Products: []int{1, 2, 3},
	},
	{
		ID:       1,
		Name:     "Samsung",
		Products: []int{4, 5},
	},
}

func GetAllProducts() []product.Product {
	return products
}

func GetAllMerchants() []merchant.Merchant {
	return merchants
}
