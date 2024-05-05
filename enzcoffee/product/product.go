package product

import (
    "database/sql"
    //"fmt"
    _ "github.com/lib/pq"
)

// 商品情報を格納する構造体
type Product struct {
	ID    int
	Name  string
	Price float64
}

func connectDB() (*sql.DB, error) {
	connStr := "user=username dbname=yourdbname password=yourpassword host=yourhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateProduct(name string, price float64) (Product, error) {
    db, err := connectDB()
    if err != nil {
        return Product{}, err
    }
    defer db.Close()

    query := `INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id`
    var newProduct Product
    newProduct.Name = name
    newProduct.Price = price
    err = db.QueryRow(query, name, price).Scan(&newProduct.ID)
    if err != nil {
        return Product{}, err
    }
    return newProduct, nil
}

func GetProductInfo(id int) (Product, error) {
	db, err := connectDB()
	if err != nil {
		return Product{}, err
	}
	defer db.Close()

	var product Product
	query := `SELECT id, name, price FROM products WHERE id = $1`
	err = db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func UpdateProductInfo(id int, newName string, newPrice float64) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `UPDATE products SET name = $2, price = $3 WHERE id = $1`
	_, err = db.Exec(query, id, newName, newPrice)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProductInfo(id int) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `DELETE FROM products WHERE id = $1`
	_, err = db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
