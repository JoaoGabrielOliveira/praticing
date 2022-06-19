namespace Models {
    public abstract class Model {
        public object Id;

        protected Model(object id)
        {
            Id = id;
        }
    }
}