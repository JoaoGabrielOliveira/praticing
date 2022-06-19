using Models;

namespace Repositories
{
    public interface IRepository<T> where T : Model  {
        T GetBy(string param, object value);
        T GetById(object id);

        List<T> All();

        void Add(T model);
        void Update(T model);
        void Delete(T model);
    }
}