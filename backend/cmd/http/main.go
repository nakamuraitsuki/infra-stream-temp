package main

import "log"

func main() {
	e, err := InitializeHTTPServer()
	if err != nil {
		log.Fatalf("サーバーの初期化に失敗しました: %v", err)
	}
	
	e.Logger.Fatal(e.Start(":8080"))
}
