package services

type CartProducts struct {
	Id        int `json:"id"`
	CartId    int `json:"cart_id"`
	ProductId int `json:"prduct_id"`
}
