from pymongo import MongoClient

# Conexión a MongoDB
cliente = MongoClient("mongodb://localhost:27017")

# Seleccionar base de datos
bd = cliente["mi_basedatos"]

# Listar colecciones (equivalente a tablas)
colecciones = bd.list_collection_names()

print("📦 Colecciones en la base de datos:")
for coleccion in colecciones:
    print("-", coleccion)

# Cerrar conexión (opcional, ya que MongoClient lo gestiona automáticamente)
cliente.close()
