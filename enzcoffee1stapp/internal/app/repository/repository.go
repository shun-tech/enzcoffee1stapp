package repository

import (
    "database/sql"
    "enzcoffee1stapp/internal/domain"
)

// db はデータベース接続を表します。実際のデータベース接続の詳細は省略します。
var db *sql.DB

// GetAllProducts はデータベースからすべての商品情報を取得します。
func GetAllProducts() ([]domain.Product, error) {
    var products []domain.Product
    rows, err := db.Query("SELECT id, name, price FROM products ORDER BY id")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var p domain.Product
        if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
            return nil, err
        }
        products = append(products, p)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return products, nil
}

// AddProduct はデータベースに新しい商品を追加します。
func AddProduct(name string, price float64) (int, error) {
    var productID int
    err := db.QueryRow("INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id", name, price).Scan(&productID)
    if err != nil {
        return 0, err
    }
    return productID, nil
}

// GetProduct は指定されたIDの商品情報をデータベースから取得します。
func GetProduct(id int) (*domain.Product, error) {
    var p domain.Product
    err := db.QueryRow("SELECT id, name, price FROM products WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.Price)
    if err != nil {
        return nil, err
    }
    return &p, nil
}

// UpdateProduct は指定されたIDの商品情報を更新します。
func UpdateProduct(p domain.Product) error {
    _, err := db.Exec("UPDATE products SET name = $1, price = $2 WHERE id = $3", p.Name, p.Price, p.ID)
    return err
}

// DeleteProduct は指定されたIDの商品をデータベースから削除します。
func DeleteProduct(id int) error {
    _, err := db.Exec("DELETE FROM products WHERE id = $1", id)
    return err
}
