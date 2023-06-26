using Microsoft.AspNetCore.Mvc;
using Models;
using Repositories;

namespace Controllers
{
    public class StudentController : ApiController<Student, long>
    {
        public StudentController() : base(new StudentRepository()) {}

        // GET: /student
        public override IEnumerable<Student> Index()
        {
            return this.Repository.Find();
        }

        // GET: /student/{id}
        public override Student Get(long id)
        {
            return this.Repository.GetById(id);
        }
    }
}