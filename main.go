package main

import (
	"API/internal/service"
	"API/internal/store"
	"API/internal/transport"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main(){
	// Conectar a SQLLite
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	//crear la tabla si no existe
	q := `
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			author TEXT NOT NULL
		)
	`
	if _, err := db.Exec(q); err != nil{
		log.Fatal(err.Error())
	}


	// Inyectar nuestras  dependencias
	bookStore := store.New(db)
	bookService := service.New(bookStore)
	BookHandler := transport.New(bookService)

	// Configurar rutas
	http.HandleFunc("/books", BookHandler.HandleBooks)
	http.HandleFunc("/books/", BookHandler.HandleBOokByID)

fmt.Println("🚀 Servidor ejecutándose en http://localhost:8080")
fmt.Println("📚 API Endpoints:")
fmt.Println("  GET     /books        - Obtener todos los libros")
fmt.Println("  POST    /books        - Crear un nuevo libro")
fmt.Println("  GET     /books/{id}   - Obtener un libro específico")
fmt.Println("  PUT     /books/{id}   - Actualizar un libro")
fmt.Println("  DELETE  /books/{id}   - Eliminar un libro")
	//Empezar y escuchar servidor
log.Fatal(http.ListenAndServe(":8080",nil))
}