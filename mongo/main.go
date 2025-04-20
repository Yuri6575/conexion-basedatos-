package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Crear contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// URI de conexión
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Conectar al servidor MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("❌ Error al conectar a MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	// Verificar conexión
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("❌ No se pudo conectar a MongoDB: %v", err)
	}

	fmt.Println("✅ Conexión exitosa a MongoDB - Base de datos 'mi_basedatos'")

	// Seleccionar base de datos
	db := client.Database("mi_basedatos")

	// Listar colecciones
	colecciones, err := db.ListCollectionNames(ctx, map[string]interface{}{})
	if err != nil {
		log.Fatalf("❌ Error al obtener colecciones: %v", err)
	}

	fmt.Println("📦 Colecciones en la base de datos:")
	for _, nombre := range colecciones {
		fmt.Println("- " + nombre)
	}
}
:= "host=localhost port=5432 user=postgres password=1234 dbname=mi_basedatos sslmode=disable"

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
