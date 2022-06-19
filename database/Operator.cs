using System.Data.SqlClient;

namespace Database;

public class Operator {
    private static SqlDataReader ExecuteCommand(SqlCommand command)
    {
        return command.ExecuteReader();
    }

    public static SqlDataReader ExecuteCommand(String query, params SqlParameter[] queryParams)
    {
        Connector connector = new Connector();
        using(SqlCommand command = new SqlCommand(query, connector.Connection)){
            command.Parameters.AddRange(queryParams);
            return command.ExecuteReader();
        }

    }

}