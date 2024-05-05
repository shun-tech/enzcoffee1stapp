package service

import (
    "enzcoffee1stapp/internal/domain"
    "enzcoffee1stapp/internal/app/repository"
)

// GetProductByID は指定されたIDの商品情報を取得します。
func GetProductByID(productID int) (*domain.Product, error) {
    return repository.GetProduct(productID)
}

// GetAllProducts はすべての商品情報を取得します。
func GetAllProducts() ([]domain.Product, error) {
    return repository.GetAllProducts()
}

// AddNewProduct は新しい商品をデータベースに追加します。
func AddNewProduct(name string, price float64) (int, error) {
    return repository.AddProduct(name, price)
}

// // UpdateProductDetails は指定されたIDの商品情報を更新します。
// func UpdateProductDetails(productID int, name string, price float64) error {
//     return repository.UpdateProduct(productID, name, price)
// }

// UpdateProduct は商品情報を更新します。
func UpdateProduct(id int, name string, price float64) error {
    // domain.Product 型のオブジェクトを作成
    product := domain.Product{
        ID:    id,
        Name:  name,
        Price: price,
    }

    // repository の UpdateProduct 関数を呼び出し
    return repository.UpdateProduct(product)
}

// DeleteProductByID は指定されたIDの商品をデータベースから削除します。
func DeleteProductByID(productID int) error {
    return repository.DeleteProduct(productID)
}