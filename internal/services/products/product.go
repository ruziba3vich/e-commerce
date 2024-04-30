package services

type Product struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Price           int    `json:"price"`
	NumberOfProduct int    `json:"number_of_product"`
}

type ProductWithNumberOfProducts struct {
	Product         Product
	NumberOfProduct int
}
