import com.mongodb.client.MongoClients;
import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoDatabase;
import org.bson.Document;

public class ConexionMongoDB {
    public static void main(String[] args) {
        // URL de conexión a MongoDB
        String uri = "mongodb://localhost:27017";

        try (MongoClient mongoClient = MongoClients.create(uri)) {
            MongoDatabase database = mongoClient.getDatabase("mi_basedatos");
            System.out.println("✅ Conexión exitosa a MongoDB");

            // Listar colecciones (equivalente a "tablas")
            System.out.println("📦 Colecciones en la base de datos:");
            for (String nombreColeccion : database.listCollectionNames()) {
                System.out.println("- " + nombreColeccion);
            }
        } catch (Exception e) {
            System.out.println("❌ Error durante la conexión: " + e.getMessage());
        }
    }
}
