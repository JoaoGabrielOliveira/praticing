using System.Data;
using System.Data.SqlClient;
namespace Database;


public class Connector : IDisposable {
    protected const string ConnectionString = "Server=localhost:1433;User Id=sa;Password=Pass1234;";
    public SqlConnection Connection;

    public Connector(){
        Connection = new SqlConnection(ConnectionString);
    }

    public void Dispose()
    {
        Connection.Close();
    }

    public void OpenConnection(){
        Connection.Open();
    }
}