using Models;
using System.Data;
using System.Data.SqlClient;

namespace Repositories {
    public class StudentRepository : IRepository<Student> {

        public Student GetBy(string param, object value)
        {
            String query = "SELECT * FROM Students WHERE @PARAM = @VALUE";
            var reader = Database.Operator.ExecuteCommand(query, new SqlParameter("PARAM", value), new SqlParameter("VALUE", value));
            
            reader.Read();
            return new Student(
                reader.GetInt32(0),
                reader.GetString(1),
                reader.GetString(2),
                reader.GetString(3),
                reader.GetFieldValue<List<Course>>(4)
            );
        }

        public Student GetById(object id)
        {
            return GetBy("id", id);
        }

        public List<Student> All()
        {
            String query = "SELECT * FROM Students";
            var reader = Database.Operator.ExecuteCommand(query);

            List<Student> students = new List<Student>();
            
            while(reader.Read())
            {
                students.Add(
                    new Student(
                        reader.GetInt32(0),
                        reader.GetString(1),
                        reader.GetString(2),
                        reader.GetString(3),
                        reader.GetFieldValue<List<Course>>(4)
                    )
                );
            }

            return students;
        }

        public void Add(Student model)
        {
            throw new NotImplementedException();
        }

        public void Update(Student model)
        {
            throw new NotImplementedException();
        }

        public void Delete(Student model)
        {
            throw new NotImplementedException();
        }
    }
}