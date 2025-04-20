const mysql = require('mysql2');

const db = mysql.createConnection({
  host: 'localhost',
  user: 'root',
  password: '',         // pon tu contraseña si tienes
  database: 'conexpython'  // asegúrate que la base de datos exista
});

db.connect((err) => {
  if (err) {
    console.error('❌ Error al conectar a MySQL:', err);
  } else {
    console.log('✅ Conexión exitosa a MySQL');
  }
});
