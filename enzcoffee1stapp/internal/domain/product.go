package domain

import (
    "database/sql"
    "time"
    "enzcoffee1stapp/pkg/db"
)

type Product struct {
    ID        int
    Name      string
    Price     float64
    CreatedAt time.Time
}

// AddProduct は新しい商品をデータベースに追加します。
func AddProduct(name string, price float64) (int, error) {
    var productID int
    createdAt := time.Now()

    query := `INSERT INTO products (name, price, created_at) VALUES ($1, $2, $3) RETURNING id`
    err := db.DB.QueryRow(query, name, price, createdAt).Scan(&productID)
    if err != nil {
        return 0, err
    }
    return productID, nil
}

// GetProduct は指定されたIDの商品情報を取得します。
func GetProduct(id int) (*Product, error) {
    var p Product
    query := `SELECT id, name, price, created_at FROM products WHERE id = $1`
    row := db.DB.QueryRow(query, id)
    err := row.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &p, nil
}

// UpdateProduct は指定されたIDの商品情報を更新します。
func UpdateProduct(id int, name string, price float64) error {
    query := `UPDATE products SET name = $2, price = $3 WHERE id = $1`
    _, err := db.DB.Exec(query, id, name, price)
    return err
}

// DeleteProduct は指定されたIDの商品をデータベースから削除します。
func DeleteProduct(id int) error {
    query := `DELETE FROM products WHERE id = $1`
    _, err := db.DB.Exec(query, id)
    return err
}
