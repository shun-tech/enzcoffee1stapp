package main

import (
	"github.com/shuntamorioka/enzcoffee/product"
	"fmt"
)

func main() {
    // product.go内のCRUDメソッドをテストする例
    testCreate()
    testRead()
    testUpdate()
    testDelete()
}

// Create操作のテスト
func testCreate() {
    // 新しい商品の名前と価格
    newName := "エスプレッソマシン"
    newPrice := 25000.00

    // CreateProductメソッドを呼び出し
    newProduct, err := product.CreateProduct(newName, newPrice)
    if err != nil {
        fmt.Printf("Createのテスト失敗: %v\n", err)
    } else {
        fmt.Printf("Createのテスト成功: 新しい商品ID %d, 商品名=%s, 価格=%.2f円\n", newProduct.ID, newProduct.Name, newProduct.Price)
    }
}

// Read操作のテスト
func testRead() {
    // テスト対象の商品ID
    testID := 1

    // GetProductInfoメソッドを呼び出し
    prod, err := product.GetProductInfo(testID)
    if err != nil {
        fmt.Printf("Readのテスト失敗: %v\n", err)
    } else {
        fmt.Printf("Readのテスト成功: 商品ID %d の商品名=%s, 価格=%.2f円\n", prod.ID, prod.Name, prod.Price)
    }
}

// Update操作のテスト
func testUpdate() {
    // 更新対象の商品ID
    updateID := 1
    // 新しい商品情報
    newName := "アップデートされたコーヒーマグ"
    newPrice := 1500.00

    // UpdateProductInfoメソッドを呼び出し
    err := product.UpdateProductInfo(updateID, newName, newPrice)
    if err != nil {
        fmt.Printf("Updateのテスト失敗: %v\n", err)
    } else {
        fmt.Println("Updateのテスト成功")
    }
}

// Delete操作のテスト
func testDelete() {
    // 削除対象の商品ID
    deleteID := 1

    // DeleteProductInfoメソッドを呼び出し
    err := product.DeleteProductInfo(deleteID)
    if err != nil {
        fmt.Printf("Deleteのテスト失敗: %v\n", err)
    } else {
        fmt.Println("Deleteのテスト成功")
    }
}