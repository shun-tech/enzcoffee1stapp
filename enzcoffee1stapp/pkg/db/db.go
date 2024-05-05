package db

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // PostgreSQLドライバー
)

var DB *sql.DB

// InitDB はデータベース接続を初期化します。
func InitDB(dataSourceName string) {
    var err error
    DB, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Fatalf("データベース接続の開設に失敗しました: %v", err)
    }

    // データベースへの接続を確認
    err = DB.Ping()
    if err != nil {
        log.Fatalf("データベースへの接続確認に失敗しました: %v", err)
    }

    fmt.Println("データベース接続が正常に確立されました。")
}

// CloseDB はデータベース接続を閉じます。
func CloseDB() {
    if err := DB.Close(); err != nil {
        log.Fatalf("データベース接続のクローズに失敗しました: %v", err)
    }
}