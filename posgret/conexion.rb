require 'pg'

begin
  # Crear la conexión a PostgreSQL
  conn = PG.connect(
    host: "localhost",
    dbname: "mi_basedatos",
    user: "postgres",
    password: "1234" # reemplaza con tu contraseña real
  )

  puts "✅ Conexión exitosa a la base de datos conex3066478"

  # Hacer una consulta
  res = conn.exec("SELECT * FROM productos")

  # Mostrar los resultados
  res.each do |fila|
    puts fila
  end

rescue PG::Error => e
  puts "❌ Error al conectar o consultar: #{e.message}"
ensure
  conn.close if conn
end
