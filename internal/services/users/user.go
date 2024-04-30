package services

import (
	"database/sql"
	"fmt"
	"sync"

	cart "github.com/ruzba3vich/e_commerce/internal/services/carts"
	products "github.com/ruzba3vich/e_commerce/internal/services/products"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Balance  int    `json:"balance"`
	Country  string `json:"country"`
	mutex    sync.Mutex
}

func (u User) AddProductIntoCart(c *cart.Cart, p products.Product, db *sql.DB) error {
	query := `
        INSERT INTO cart_products(cart_id, product_id)
            VALUES($1, $2) RETURNING id;
    `
	var cartProductId int
	err := db.QueryRow(query, c.Id, p.Id).Scan(&cartProductId)
	if err != nil {
		return err
	}

	cartQuery := `
        INSERT INTO carts(id, user_id) 
        SELECT $1, $2 
        WHERE NOT EXISTS (SELECT 1 FROM carts WHERE id = $1) 
        RETURNING id, user_id;
    `
	var updatedCart cart.Cart
	if err := db.QueryRow(cartQuery, c.Id, u.Id).Scan(&updatedCart.Id, &updatedCart.UserId); err != nil {
		return err
	}
	return nil
}

func (u User) RemoveProductFromCart(c cart.Cart, p products.Product, db *sql.DB) error {
	query := `
        DELETE FROM cart_products
        WHERE cart_id = $1 AND product_id = $2;
    `
	_, err := db.Exec(query, c.Id, p.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BuyProducts(products []products.ProductWithNumberOfProducts, db *sql.DB) error {
	for _, product := range products {
		remaining := product.NumberOfProduct
		if remaining < 0 {
			return fmt.Errorf("not enough products available for product with ID %d", product.Product.Id)
		}
		u.mutex.Lock()
		result, err := func() (sql.Result, error) {
			defer u.mutex.Unlock()
			innerResult, innerErr := db.Exec("UPDATE Products SET number_of_product = number_of_product - $1 WHERE id = $2", remaining, product.Product.Id)
			return innerResult, innerErr
		}()

		if err != nil {
			return err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return fmt.Errorf("product with ID %d not found", product.Product.Id)
		}
	}
	return nil
}
