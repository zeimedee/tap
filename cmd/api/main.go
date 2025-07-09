package main

import (
	"os"

	"github.com/zeimedee/loverboy/internal/routes"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("error loading env")
	// }

	port := os.Getenv("PORT")
	router := routes.SetUpRoutes()

	router.Run(port)

}
