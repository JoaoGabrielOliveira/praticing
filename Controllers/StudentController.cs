using Microsoft.AspNetCore.Mvc;
using Models;
using Repositories;

namespace Controllers
{
    public class StudentController : Controller
    {
        // GET: /student
        public IEnumerable<Student> Index()
        {
            List<Student> students = new List<Student>()
            {
                new Student(0, "João", "Conceição", "1123123", new List<Course>())
            };

            return students;
        }

        // GET: /student/{id}
        public Student Get(int id)
        {
            return new Student(id, "João", "Conceição", "1123123", new List<Course>());
        }
    }
}