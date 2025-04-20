#include <iostream>
#include <stdexcept>
#include <mysql_driver.h>
#include <mysql_connection.h>

using namespace std;

int main() {
    try {
        // Crear el driver y la conexión
        sql::mysql::MySQL_Driver *driver;
        sql::Connection *con;

        // Crear el driver
        driver = sql::mysql::get_mysql_driver_instance();
        
        // Establecer la conexión a la base de datos MySQL
        con = driver->connect("tcp://127.0.0.1:3306", "root", "");  // Ajusta el usuario y contraseña
        con->setSchema("conex3066478");  // La base de datos que estás usando

        cout << "Conexión exitosa a la base de datos 'conex3066478'" << endl;

        // Hacer una consulta simple
        sql::Statement *stmt = con->createStatement();
        sql::ResultSet *res = stmt->executeQuery("SHOW TABLES");

        cout << "Tablas disponibles en la base de datos:" << endl;
        while (res->next()) {
            cout << res->getString(1) << endl;
        }

        // Liberar recursos
        delete res;
        delete stmt;
        delete con;
    } catch (sql::SQLException &e) {
        cout << "Error al conectar a MySQL: " << e.what() << endl;
    }

    return 0;
}
