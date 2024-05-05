package main

import (
    "log"
    "html/template"
    "net/http"
    "enzcoffee1stapp/internal/app/repository"
    "enzcoffee1stapp/pkg/db"
    "strconv"
    "os"
    "database/sql"
    _ "github.com/lib/pq" // PostgreSQLドライバをインポート
)

var database *sql.DB // dbという名前を避けるためにdatabaseという名前を使用

func init() {
    var err error
    // 標準のsql.Openを使用してデータベースに接続
    database, err = sql.Open("postgres", "host=localhost port=5432 user=enzcoffee password=enzcoffee dbname=enzcoffee sslmode=disable")
    if err != nil {
        log.Fatal("データベースのオープンに失敗しました:", err)
    }
    if err = database.Ping(); err != nil {
        log.Fatal("データベースへの接続に失敗しました:", err)
    }
}

func main() {
        // 環境変数を設定
    os.Setenv("WEB_DIR", "/Users/shuntamorioka/workspace/enzcoffee1stapp/web/templates/")
    // 環境変数からウェブディレクトリのパスを取得
    webDir := os.Getenv("WEB_DIR")

    // データベース接続の初期化
    db.InitDB("host=localhost port=5432 user=enzcoffee password=enzcoffee dbname=enzcoffee sslmode=disable")

    // 静的ファイルのサービング
    fs := http.FileServer(http.Dir(webDir + "/static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, webDir + "web/index.html")
    })
    http.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, webDir + "result/success.html")
    })
    http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, webDir + "result/error.html")
    })
    http.HandleFunc("/add-product", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
            return
        }
        name := r.FormValue("name")
        price, err := strconv.ParseFloat(r.FormValue("price"), 64)
        if err != nil {
            http.Redirect(w, r, "result/error.html", http.StatusFound)
            return
        }
        _, err = repository.AddProduct(name, price)
        if err != nil {
            http.Redirect(w, r, "result/error.html", http.StatusFound)
            return
        }
        http.Redirect(w, r, "result/success.html", http.StatusFound)
    })
    http.HandleFunc("/products", productsHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
    // 環境変数からウェブディレクトリのパスを取得
    webDir := os.Getenv("WEB_DIR")
    if webDir == "" {
        log.Fatal("WEB_DIR environment variable not set")
    }

    products, err := repository.GetAllProducts()
    if err != nil {
        log.Printf("Error retrieving products: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    tmpl, err := template.ParseFiles(webDir + "web/products.html")
    if err != nil {
        log.Printf("Error parsing template: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, products)
    if err != nil {
        log.Printf("Error executing template: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}
