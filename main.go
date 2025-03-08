package main

import (
	"demo/src/data"
	employeeDependencies "demo/src/employees/dependencies"
	prodDependcies "demo/src/products/dependencies"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	
	mysql := data.NewMySQL()
	defer mysql.Close()
	
	router := gin.Default()

	// Configuración de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	prodDependcies := prodDependcies.NewProductDependencies(mysql.DB)
	prodDependcies.Execute(router)

	employeeDependencies := employeeDependencies.NewEmployeeDependencies(mysql.DB)
	employeeDependencies.Execute(router)

	log.Println("[Main] Servidor corriendo en http://localhost:8080")
	router.Run(":8080")
}
