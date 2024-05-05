package domain

import (
    "database/sql"
    "fmt"
    "enzcoffee1stapp/pkg/db"
)

type Sale struct {
    ProductID int
    Quantity  int
    Total     float64
}

// CalculateTotal は指定された商品IDと数量に基づいて合計金額を計算します。
func CalculateTotal(productID, quantity int) (float64, error) {
    var price float64
    query := `SELECT price FROM products WHERE id = $1`
    err := db.DB.QueryRow(query, productID).Scan(&price)
    if err != nil {
        if err == sql.ErrNoRows {
            return 0, fmt.Errorf("指定された商品IDが見つかりません: %d", productID)
        }
        return 0, err
    }

    total := price * float64(quantity)
    return total, nil
}