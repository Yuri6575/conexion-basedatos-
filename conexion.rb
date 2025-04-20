require 'mysql2'

begin
  # Crear la conexión
  client = Mysql2::Client.new(
    host: "localhost",
    username: "root",           # cambia según tu usuario
    password: "",               # cambia según tu contraseña
    database: "conex3066478"    # tu base de datos
  )

  puts "✅ Conexión exitosa a la base de datos conex3066478"

  # Hacer una consulta
  results = client.query("SELECT * FROM productos")

  # Mostrar los resultados
  results.each do |fila|
    puts fila
  end

rescue Mysql2::Error => e
  puts "❌ Error al conectar: #{e.message}"
ensure
  client.close if client
end
