package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	// Cadena de conexión para PostgreSQL
	connStr := "host=localhost port=5432 user=postgres password=1234 dbname=mi_basedatos sslmode=disable"

	// Abrir conexión
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Error al abrir la conexión: %v", err)
	}
	defer db.Close()

	// Verificar conexión
	err = db.Ping()
	if err != nil {
		log.Fatalf("❌ No se pudo conectar: %v", err)
	}

	fmt.Println("✅ Conexión exitosa a PostgreSQL - Base de datos 'mi_basedatos'")

	// Mostrar tablas del esquema 'public'
	rows, err := db.Query(`
		SELECT table_name 
		FROM information_schema.tables 
		WHERE table_schema = 'public'
	`)
	if err != nil {
		log.Fatalf("❌ Error en la consulta: %v", err)
	}
	defer rows.Close()

	fmt.Println("📦 Tablas en la base de datos:")
	var tableName string
	for rows.Next() {
		rows.Scan(&tableName)
		fmt.Println("- " + tableName)
	}
}
