const { MongoClient } = require('mongodb');

// URL de conexión
const uri = "mongodb://localhost:27017";

// Crear cliente
const client = new MongoClient(uri);

async function conectarMongo() {
  try {
    await client.connect();
    console.log("✅ Conexión exitosa a MongoDB");

    // Seleccionar base de datos
    const db = client.db("mi_basedatos");

    // Listar colecciones
    const colecciones = await db.listCollections().toArray();
    console.log("📦 Colecciones en la base de datos:");
    colecciones.forEach(col => console.log("- " + col.name));

  } catch (error) {
    console.error("❌ Error al conectar a MongoDB:", error);
  } finally {
    await client.close();
  }
}

conectarMongo();

