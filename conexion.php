<?php
$host = "localhost";     // o la IP del servidor
$usuario = "root"; // cambia por tu usuario
$clave = "";// cambia por tu contraseña
$bd = "conex3066478";       // cambia por el nombre de tu base de datos

// Crear conexión
$conexion = new mysqli($host, $usuario, $clave, $bd);

// Verificar conexión
if ($conexion->connect_error) {
    die("Conexión fallida: " . $conexion->connect_error);
} else {
    echo "Conexión exitosa a la base de datos";
}
?>
