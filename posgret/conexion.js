const { Client } = require('pg');

const client = new Client({
  host: 'localhost',
  user: 'postgres',
  password: '1234',  // reemplaza con tu contraseña real
  database: 'mi_basedatos',
  port: 5432, // puerto por defecto de PostgreSQL
});

client.connect(err => {
  if (err) {
    console.error('❌ Error al conectar a PostgreSQL:', err.stack);
  } else {
    console.log('✅ Conexión exitosa a PostgreSQL');
  }
});

