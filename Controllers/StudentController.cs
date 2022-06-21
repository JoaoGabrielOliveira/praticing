using Microsoft.AspNetCore.Mvc;
using Models;
using Repositories;

namespace Controllers
{
    public class StudentController : ApiController<Student>
    {
        public StudentController() : base(new StudentRepository()) {}

        // GET: /student
        public override IEnumerable<Student> Index()
        {
            List<Student> students = new List<Student>()
            {
                new Student(0, "João", "Conceição", "1123123", new List<Course>())
            };

            return students;
        }

        // GET: /student/{id}
        public override Student Get(object id)
        {
            return new Student((int)id, "João", "Conceição", "1123123", new List<Course>());
        }
    }
}