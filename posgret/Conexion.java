import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;

public class ConexionPostgreSQL {
    public static void main(String[] args) {
        String url = "jdbc:postgresql://localhost:5432/mi_basedatos";
        String usuario = "postgres";
        String contrasena = "1234"; // Cambia esto por tu contraseña real

        try {
            Connection conn = DriverManager.getConnection(url, usuario, contrasena);
            System.out.println("✅ Conexión exitosa a PostgreSQL.");
            conn.close();
        } catch (SQLException e) {
            System.out.println("❌ Error durante la conexión: " + e.getMessage());
        }
    }
}
