package handler

import (
    "encoding/json"
    "net/http"
    "strconv"
    "enzcoffee1stapp/internal/app/service"
    "enzcoffee1stapp/internal/domain"
)

// HandleProducts は商品に関するCRUD操作を行うハンドラーです。
func HandleProducts(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        // 商品一覧を取得
        products, err := service.GetAllProducts()
        if err != nil {
            http.Error(w, "商品の取得に失敗しました: "+err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(products)
    case "POST":
        // 新しい商品を追加
        var product domain.Product
        if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
            http.Error(w, "リクエストの解析に失敗しました: "+err.Error(), http.StatusBadRequest)
            return
        }
        id, err := service.AddNewProduct(product.Name, product.Price)
        if err != nil {
            http.Error(w, "商品の追加に失敗しました: "+err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]int{"id": id})
    default:
        http.Error(w, "許可されていないメソッドです", http.StatusMethodNotAllowed)
    }
}

// HandleCalculateTotal は商品IDと数量を受け取り、合計金額を計算して返します。
func HandleCalculateTotal(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "POSTメソッドのみ許可されています", http.StatusMethodNotAllowed)
        return
    }

    productID, err := strconv.Atoi(r.FormValue("productID"))
    if err != nil {
        http.Error(w, "商品IDが無効です: "+err.Error(), http.StatusBadRequest)
        return
    }

    quantity, err := strconv.Atoi(r.FormValue("quantity"))
    if err != nil {
        http.Error(w, "数量が無効です: "+err.Error(), http.StatusBadRequest)
        return
    }

    total, err := domain.CalculateTotal(productID, quantity)
    if err != nil {
        http.Error(w, "合計金額の計算中にエラーが発生しました: "+err.Error(), http.StatusInternalServerError)
        return
    }

    response := struct {
        Total float64 `json:"total"`
    }{
        Total: total,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}