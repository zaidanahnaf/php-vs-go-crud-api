package main

import (
	"fmt"
	"umkm-api-go/config"
	"umkm-api-go/routes"
)

func main() {
	config.ConnectDB()

	r := routes.SetupRouter()
	fmt.Println("Server berjalan di http://localhost:8080")
	r.Run(":8080")
}
