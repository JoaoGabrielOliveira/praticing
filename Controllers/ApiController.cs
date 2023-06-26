using Repositories;

namespace Controllers
{
    public abstract class ApiController<T, ID> : Microsoft.AspNetCore.Mvc.Controller where T : Models.Model<ID>
    {
        protected IRepository<T, ID> Repository;

        public ApiController(IRepository<T, ID> repository)
        {
            Repository = repository;
        }

        public virtual IEnumerable<T> Index()
        {
            return Repository.Find();
        }
        public virtual T Get(ID id)
        {
            return Repository.GetById(id);
        }

    }
    
}