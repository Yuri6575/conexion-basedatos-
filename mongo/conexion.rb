require 'mongo'

begin
  # Crear el cliente de MongoDB
  cliente = Mongo::Client.new(['127.0.0.1:27017'], database: 'mi_basedatos')

  puts "✅ Conexión exitosa a la base de datos mi_basedatos"

  # Seleccionar la colección (como una tabla)
  coleccion = cliente[:productos]

  # Consultar todos los documentos
  coleccion.find.each do |documento|
    puts documento
  end

rescue => e
  puts "❌ Error al conectar o consultar: #{e.message}"
ensure
  cliente&.close
end
