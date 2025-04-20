#include <iostream>
#include <mongocxx/client.hpp>
#include <mongocxx/instance.hpp>
#include <mongocxx/uri.hpp>
#include <bsoncxx/json.hpp>

int main()
{
    try
    {
        // Inicializar MongoDB
        mongocxx::instance inst{};
        mongocxx::client conn{mongocxx::uri{"mongodb://localhost:27017"}};

        // Conectar a base de datos
        auto db = conn["mi_basedatos"];
        std::cout << "✅ Conexión exitosa a MongoDB\n";

        // Listar colecciones
        std::cout << "📦 Colecciones en la base de datos:" << std::endl;
        for (auto nombre : db.list_collection_names())
        {
            std::cout << "- " << nombre << std::endl;
        }
    }
    catch (const std::exception &e)
    {
        std::cerr << "❌ Error: " << e.what() << std::endl;
    }

    return 0;
}
