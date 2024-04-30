package services

import (
	"database/sql"
	"errors"

	services "github.com/ruzba3vich/e_commerce/internal/services/products"
)

type Admin struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (a Admin) AddNewProduct(p services.Product, db *sql.DB) (services.Product, error) {
	query := "INSERT INTO Products(name, price, number_of_product) VALUES ($1, $2, $3) RETURNING id;"
	err := db.QueryRow(query, p.Name, p.Price, p.NumberOfProduct).Scan(&p.Id)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (a Admin) DeleteProduct(p services.Product, db *sql.DB) (services.Product, error) {
	query := "DELETE FROM Products WHERE id = $1 RETURNING id, name"

	var deletedProduct services.Product
	err := db.QueryRow(query, p.Id).Scan(&deletedProduct.Id, &deletedProduct.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return services.Product{}, errors.New("product not found")
		}
		return services.Product{}, err
	}

	return deletedProduct, nil
}

func (a Admin) UpdateProduct(p services.Product, db *sql.DB) (services.Product, error) {
	query := `
        UPDATE products
        SET name = $1, price = $2, number_of_product = $3
        WHERE id = $4 RETURNING name, price, number_of_product;
    `
	var updatedProduct services.Product
	err := db.QueryRow(query, p.Name, p.Price, p.NumberOfProduct, p.Id).Scan(&updatedProduct.Name, &updatedProduct.Price, &updatedProduct.NumberOfProduct)

	if err != nil {
		return p, err
	}
	updatedProduct.Id = p.Id
	return updatedProduct, nil
}

func GetAllUsersByMonth()
