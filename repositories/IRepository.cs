using Models;
using Database;

namespace Repositories
{
    public interface IRepository<T, ID> where T : Model<ID>  {
        T GetById(ID id);

        List<T> Find();

        void Add(T model);
        void Update(T model);
        void Delete(T model);
    }
}