package main

import "log"

func main() {
	app, err := InitializeHTTPServer()
	if err != nil {
		log.Fatalf("サーバーの初期化に失敗しました: %v", err)
	}
	defer app.DB.Close()
	
	app.Echo.Logger.Fatal(app.Echo.Start(":8080"))
}
