#include <iostream>
#include <pqxx/pqxx>

int main()
{
    try
    {
        pqxx::connection conn("host=localhost port=5432 dbname=mi_basedatos user=postgres password=1234");

        if (conn.is_open())
        {
            std::cout << "✅ Conexión exitosa a PostgreSQL" << std::endl;

            pqxx::work txn(conn);
            pqxx::result res = txn.exec("SELECT table_name FROM information_schema.tables WHERE table_schema='public'");

            std::cout << "📋 Tablas en la base de datos:" << std::endl;
            for (auto row : res)
            {
                std::cout << "- " << row[0].c_str() << std::endl;
            }

            conn.disconnect();
        }
        else
        {
            std::cout << "❌ No se pudo abrir la conexión a PostgreSQL" << std::endl;
        }
    }
    catch (const std::exception &e)
    {
        std::cerr << "❌ Error: " << e.what() << std::endl;
    }

    return 0;
}
