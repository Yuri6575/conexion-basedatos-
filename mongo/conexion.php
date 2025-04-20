<?php
$host = "localhost";
$puerto = "5432";
$usuario = "postgres"; // tu usuario de PostgreSQL
$clave = "1234"; 
$bd = "mi_basedatos"; 

// Cadena de conexión
$conexion = pg_connect("host=$host port=$puerto dbname=$bd user=$usuario password=$clave");

// Verificar conexión
if (!$conexion) {
    echo "❌ Error al conectar a PostgreSQL";
} else {
    echo "✅ Conexión exitosa a PostgreSQL";
}
?>