package main

import (
	"database/sql"
	db2 "github.com/gui-meireles/go-hexagonal/adapters/db"
	"github.com/gui-meireles/go-hexagonal/application"
)

func main() {
	// Cria a conexão com o banco por meio do sqlite3
	db, _ := sql.Open("sqlite3", "sqlite.db")

	//  Instancia o adapter do banco (select,insert,update)
	productDbAdapter := db2.NewProductDb(db)

	// Chama a application passando o adapter do banco para fazer a persistencia
	productService := application.NewProductService(productDbAdapter)

	// Cria um produto chamando a application
	product, _ := productService.Create("Abacaxi", 30)

	// Chama o método de update, pois o produto existe
	productService.Enable(product)
}
