import psycopg2

# Conexión a PostgreSQL
conexion = psycopg2.connect(
    host="localhost",
    user="postgres",
    password="1234",
    dbname="mi_basedatos"
)

cursor = conexion.cursor()

# Mostrar tablas (esquema público por defecto)
cursor.execute("""
    SELECT table_name 
    FROM information_schema.tables 
    WHERE table_schema = 'public'
""")

# Mostrar nombres de tablas
for tabla in cursor.fetchall():
    print(tabla)

cursor.close()
conexion.close()
