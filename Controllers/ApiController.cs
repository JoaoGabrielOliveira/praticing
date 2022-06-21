using Repositories;

namespace Controllers
{
    public abstract class ApiController<T> : Microsoft.AspNetCore.Mvc.Controller where T : Models.Model
    {
        protected IRepository<T> Repository;

        public ApiController(IRepository<T> repository)
        {
            Repository = repository;
        }

        public virtual IEnumerable<T> Index()
        {
            return Repository.All();
        }
        public virtual T Get(object id)
        {
            return Repository.GetById(id);
        }

    }
    
}