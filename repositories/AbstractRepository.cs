using Models;
using Database;

namespace Repositories {

    public abstract class AbstractRepository<T, ID> : IRepository<T, ID> where T : Model<ID>
    {
        private BaseContext<T, ID> Context { get; set; }

        public AbstractRepository()
        {
            Context = new BaseContext<T, ID>();
        }

        public T GetById(ID id) => Context.Data.Find(id);

        public List<T> Find() => Context.Data.OrderBy(model => model.Id).ToList();

        public void Add(T model)
        {
            Context.Add(model);
            Context.SaveChanges();
        }

        public void Update(T model)
        {
            Context.Update(model);
            Context.SaveChanges();
        }

        public void Delete(T model)
        {
            Context.Remove(model);
            Context.SaveChanges();
        }
    }
}